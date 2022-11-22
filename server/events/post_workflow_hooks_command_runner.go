package events

import (
	"github.com/runatlantis/atlantis/server/core/config/valid"
	"github.com/runatlantis/atlantis/server/core/runtime"
	"github.com/runatlantis/atlantis/server/events/command"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/events/vcs"
)

//go:generate mockgen -package mocks -destination mocks/mock_post_workflow_hooks_command_runner.go . PostWorkflowHooksCommandRunner

type PostWorkflowHooksCommandRunner interface {
	RunPostHooks(ctx *command.Context, cmd *CommentCommand) error
}

// DefaultPostWorkflowHooksCommandRunner is the first step when processing a workflow hook commands.
type DefaultPostWorkflowHooksCommandRunner struct {
	VCSClient              vcs.Client
	WorkingDirLocker       WorkingDirLocker
	WorkingDir             WorkingDir
	GlobalCfg              valid.GlobalCfg
	PostWorkflowHookRunner runtime.PostWorkflowHookRunner
}

// RunPostHooks runs post_workflow_hooks after a plan/apply has completed
func (w *DefaultPostWorkflowHooksCommandRunner) RunPostHooks(
	ctx *command.Context, cmd *CommentCommand,
) error {
	pull := ctx.Pull
	baseRepo := pull.BaseRepo
	headRepo := ctx.HeadRepo
	user := ctx.User
	log := ctx.Log

	postWorkflowHooks := make([]*valid.WorkflowHook, 0)
	for _, repo := range w.GlobalCfg.Repos {
		if repo.IDMatches(baseRepo.ID()) && repo.BranchMatches(pull.BaseBranch) && len(repo.PostWorkflowHooks) > 0 {
			postWorkflowHooks = append(postWorkflowHooks, repo.PostWorkflowHooks...)
		}
	}

	// short circuit any other calls if there are no post-hooks configured
	if len(postWorkflowHooks) == 0 {
		return nil
	}

	log.Debug("post-hooks configured, running...")

	unlockFn, err := w.WorkingDirLocker.TryLock(baseRepo.FullName, pull.Num, DefaultWorkspace, DefaultRepoRelDir)
	if err != nil {
		return err
	}
	log.Debug("got workspace lock")
	defer unlockFn()

	repoDir, _, err := w.WorkingDir.Clone(log, headRepo, pull, DefaultWorkspace)
	if err != nil {
		return err
	}

	var escapedArgs []string
	if cmd != nil {
		escapedArgs = escapeArgs(cmd.Flags)
	}

	err = w.runHooks(
		models.WorkflowHookCommandContext{
			BaseRepo:           baseRepo,
			HeadRepo:           headRepo,
			Log:                log,
			Pull:               pull,
			User:               user,
			Verbose:            false,
			EscapedCommentArgs: escapedArgs,
		},
		postWorkflowHooks, repoDir)

	if err != nil {
		return err
	}

	return nil
}

func (w *DefaultPostWorkflowHooksCommandRunner) runHooks(
	ctx models.WorkflowHookCommandContext,
	postWorkflowHooks []*valid.WorkflowHook,
	repoDir string,
) error {

	for _, hook := range postWorkflowHooks {
		_, err := w.PostWorkflowHookRunner.Run(ctx, hook.RunCommand, repoDir)

		if err != nil {
			return err
		}
	}

	return nil
}
