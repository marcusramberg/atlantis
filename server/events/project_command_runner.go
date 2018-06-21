// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.
//
package events

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/events/runtime"
	"github.com/runatlantis/atlantis/server/events/webhooks"
	"github.com/runatlantis/atlantis/server/events/yaml/valid"
	"github.com/runatlantis/atlantis/server/logging"
)

//go:generate pegomock generate -m --use-experimental-model-gen --package mocks -o mocks/mock_lock_url_generator.go LockURLGenerator

type LockURLGenerator interface {
	GenerateLockURL(lockID string) string
}

type WebhooksSender interface {
	Send(log *logging.SimpleLogger, result webhooks.ApplyResult) error
}

// PlanSuccess is the result of a successful plan.
type PlanSuccess struct {
	TerraformOutput string
	LockURL         string
}

type ProjectCommandRunner struct {
	Locker                  ProjectLocker
	LockURLGenerator        LockURLGenerator
	InitStepRunner          runtime.InitStepRunner
	PlanStepRunner          runtime.PlanStepRunner
	ApplyStepRunner         runtime.ApplyStepRunner
	RunStepRunner           runtime.RunStepRunner
	PullApprovedChecker     runtime.PullApprovedChecker
	Workspace               AtlantisWorkspace
	Webhooks                WebhooksSender
	AtlantisWorkspaceLocker AtlantisWorkspaceLocker
}

func (p *ProjectCommandRunner) Plan(ctx models.ProjectCommandContext) ProjectCommandResult {
	// Acquire Atlantis lock for this repo/dir/workspace.
	lockAttempt, err := p.Locker.TryLock(ctx.Log, ctx.Pull, ctx.User, ctx.Workspace, models.NewProject(ctx.BaseRepo.FullName, ctx.RepoRelPath))
	if err != nil {
		return ProjectCommandResult{
			Error: errors.Wrap(err, "acquiring lock"),
		}
	}
	if !lockAttempt.LockAcquired {
		return ProjectCommandResult{Failure: lockAttempt.LockFailureReason}
	}
	ctx.Log.Debug("acquired lock for project")

	// Acquire internal lock for the directory we're going to operate in.
	unlockFn, err := p.AtlantisWorkspaceLocker.TryLock2(ctx.BaseRepo.FullName, ctx.Workspace, ctx.Pull.Num)
	if err != nil {
		return ProjectCommandResult{Error: err}
	}
	defer unlockFn()

	// Clone is idempotent so okay to run even if the repo was already cloned.
	repoDir, cloneErr := p.Workspace.Clone(ctx.Log, ctx.BaseRepo, ctx.HeadRepo, ctx.Pull, ctx.Workspace)
	if cloneErr != nil {
		if unlockErr := lockAttempt.UnlockFn(); unlockErr != nil {
			ctx.Log.Err("error unlocking state after plan error: %v", unlockErr)
		}
		return ProjectCommandResult{Error: cloneErr}
	}
	projAbsPath := filepath.Join(repoDir, ctx.RepoRelPath)

	// Use default stage unless another workflow is defined in config
	stage := p.defaultPlanStage()
	if ctx.ProjectConfig != nil && ctx.ProjectConfig.Workflow != nil {
		ctx.Log.Debug("project configured to use workflow %q", *ctx.ProjectConfig.Workflow)
		configuredStage := ctx.GlobalConfig.GetPlanStage(*ctx.ProjectConfig.Workflow)
		if configuredStage != nil {
			ctx.Log.Debug("project will use the configured stage for that workflow")
			stage = *configuredStage
		}
	}
	outputs, err := p.runSteps(stage.Steps, ctx, projAbsPath)
	if err != nil {
		if unlockErr := lockAttempt.UnlockFn(); unlockErr != nil {
			ctx.Log.Err("error unlocking state after plan error: %v", unlockErr)
		}
		// todo: include output from other steps.
		return ProjectCommandResult{Error: err}
	}

	return ProjectCommandResult{
		PlanSuccess: &PlanSuccess{
			LockURL:         p.LockURLGenerator.GenerateLockURL(lockAttempt.LockKey),
			TerraformOutput: strings.Join(outputs, "\n"),
		},
	}
}

func (p *ProjectCommandRunner) runSteps(steps []valid.Step, ctx models.ProjectCommandContext, absPath string) ([]string, error) {
	var outputs []string
	for _, step := range steps {
		var out string
		var err error
		switch step.StepName {
		case "init":
			out, err = p.InitStepRunner.Run(ctx, step.ExtraArgs, absPath)
		case "plan":
			out, err = p.PlanStepRunner.Run(ctx, step.ExtraArgs, absPath)
		case "apply":
			out, err = p.ApplyStepRunner.Run(ctx, step.ExtraArgs, absPath)
		case "run":
			out, err = p.RunStepRunner.Run(ctx, step.RunCommand, absPath)
		}

		if err != nil {
			// todo: include output from other steps.
			return nil, err
		}
		if out != "" {
			outputs = append(outputs, out)
		}
	}
	return outputs, nil
}

func (p *ProjectCommandRunner) Apply(ctx models.ProjectCommandContext) ProjectCommandResult {
	repoDir, err := p.Workspace.GetWorkspace(ctx.BaseRepo, ctx.Pull, ctx.Workspace)
	if err != nil {
		if os.IsNotExist(err) {
			return ProjectCommandResult{Error: errors.New("project has not been cloned–did you run plan?")}
		}
		return ProjectCommandResult{Error: err}
	}
	absPath := filepath.Join(repoDir, ctx.RepoRelPath)

	if ctx.ProjectConfig != nil {
		for _, req := range ctx.ProjectConfig.ApplyRequirements {
			switch req {
			case "approved":
				approved, err := p.PullApprovedChecker.PullIsApproved(ctx.BaseRepo, ctx.Pull)
				if err != nil {
					return ProjectCommandResult{Error: errors.Wrap(err, "checking if pull request was approved")}
				}
				if !approved {
					return ProjectCommandResult{Failure: "Pull request must be approved before running apply."}
				}
			}
		}
	}
	// Acquire internal lock for the directory we're going to operate in.
	unlockFn, err := p.AtlantisWorkspaceLocker.TryLock2(ctx.BaseRepo.FullName, ctx.Workspace, ctx.Pull.Num)
	if err != nil {
		return ProjectCommandResult{Error: err}
	}
	defer unlockFn()

	// Use default stage unless another workflow is defined in config
	stage := p.defaultApplyStage()
	if ctx.ProjectConfig != nil && ctx.ProjectConfig.Workflow != nil {
		configuredStage := ctx.GlobalConfig.GetApplyStage(*ctx.ProjectConfig.Workflow)
		if configuredStage != nil {
			stage = *configuredStage
		}
	}
	outputs, err := p.runSteps(stage.Steps, ctx, absPath)
	p.Webhooks.Send(ctx.Log, webhooks.ApplyResult{ // nolint: errcheck
		Workspace: ctx.Workspace,
		User:      ctx.User,
		Repo:      ctx.BaseRepo,
		Pull:      ctx.Pull,
		Success:   err == nil,
	})
	if err != nil {
		// todo: include output from other steps.
		return ProjectCommandResult{Error: err}
	}
	return ProjectCommandResult{
		ApplySuccess: strings.Join(outputs, "\n"),
	}
}

func (p ProjectCommandRunner) defaultPlanStage() valid.Stage {
	return valid.Stage{
		Steps: []valid.Step{
			{
				StepName: "init",
			},
			{
				StepName: "plan",
			},
		},
	}
}

func (p ProjectCommandRunner) defaultApplyStage() valid.Stage {
	return valid.Stage{
		Steps: []valid.Step{
			{
				StepName: "apply",
			},
		},
	}
}
