package events_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	runtime_mocks "github.com/runatlantis/atlantis/server/core/runtime/mocks"

	"github.com/runatlantis/atlantis/server/core/config/valid"
	"github.com/runatlantis/atlantis/server/events"
	"github.com/runatlantis/atlantis/server/events/command"
	"github.com/runatlantis/atlantis/server/events/mocks"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/events/models/fixtures"
	vcsmocks "github.com/runatlantis/atlantis/server/events/vcs/mocks"
	"github.com/runatlantis/atlantis/server/logging"
	. "github.com/runatlantis/atlantis/testing"
)

var postWh events.DefaultPostWorkflowHooksCommandRunner
var postWhWorkingDir *mocks.MockWorkingDir
var postWhWorkingDirLocker *mocks.MockWorkingDirLocker
var whPostWorkflowHookRunner *runtime_mocks.MockPostWorkflowHookRunner

func postWorkflowHooksSetup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	vcsClient := vcsmocks.NewMockClient(ctrl)
	postWhWorkingDir = mocks.NewMockWorkingDir(ctrl)
	postWhWorkingDirLocker = mocks.NewMockWorkingDirLocker(ctrl)
	whPostWorkflowHookRunner = runtime_mocks.NewMockPostWorkflowHookRunner(ctrl)

	postWh = events.DefaultPostWorkflowHooksCommandRunner{
		VCSClient:              vcsClient,
		WorkingDirLocker:       postWhWorkingDirLocker,
		WorkingDir:             postWhWorkingDir,
		PostWorkflowHookRunner: whPostWorkflowHookRunner,
	}
}

func TestRunPostHooks_Clone(t *testing.T) {

	log := logging.NewNoopLogger(t)

	var newPull = fixtures.Pull
	newPull.BaseRepo = fixtures.GithubRepo

	ctx := &command.Context{
		Pull:     newPull,
		HeadRepo: fixtures.GithubRepo,
		User:     fixtures.User,
		Log:      log,
	}

	testHook := valid.WorkflowHook{
		StepName:   "test",
		RunCommand: "some command",
	}

	pCtx := models.WorkflowHookCommandContext{
		BaseRepo: fixtures.GithubRepo,
		HeadRepo: fixtures.GithubRepo,
		Pull:     newPull,
		Log:      log,
		User:     fixtures.User,
		Verbose:  false,
	}

	repoDir := "path/to/repo"
	result := "some result"

	t.Run("success hooks in cfg", func(t *testing.T) {
		postWorkflowHooksSetup(t)

		unlockCalled := newBool(false)
		unlockFn := func() {
			unlockCalled = newBool(true)
		}

		globalCfg := valid.GlobalCfg{
			Repos: []valid.Repo{
				{
					ID: fixtures.GithubRepo.ID(),
					PostWorkflowHooks: []*valid.WorkflowHook{
						&testHook,
					},
				},
			},
		}

		postWh.GlobalCfg = globalCfg

		postWhWorkingDirLocker.EXPECT().TryLock(fixtures.GithubRepo.FullName, newPull.Num, events.DefaultWorkspace, events.DefaultRepoRelDir).Return(unlockFn, nil)
		postWhWorkingDir.EXPECT().Clone(log, fixtures.GithubRepo, newPull, events.DefaultWorkspace).Return(repoDir, false, nil)
		whPostWorkflowHookRunner.EXPECT().Run(pCtx, testHook.RunCommand, repoDir).Return(result, nil)

		err := postWh.RunPostHooks(ctx, nil)

		Ok(t, err)
		whPostWorkflowHookRunner.EXPECT().Run(pCtx, testHook.RunCommand, repoDir).Times(1)
		Assert(t, *unlockCalled == true, "unlock function called")
	})
	t.Run("success hooks not in cfg", func(t *testing.T) {
		postWorkflowHooksSetup(t)
		globalCfg := valid.GlobalCfg{
			Repos: []valid.Repo{
				// one with hooks but mismatched id
				{
					ID: "id1",
					PostWorkflowHooks: []*valid.WorkflowHook{
						&testHook,
					},
				},
				// one with the correct id but no hooks
				{
					ID:                fixtures.GithubRepo.ID(),
					PostWorkflowHooks: []*valid.WorkflowHook{},
				},
			},
		}

		postWh.GlobalCfg = globalCfg

		err := postWh.RunPostHooks(ctx, nil)

		Ok(t, err)

		whPostWorkflowHookRunner.EXPECT().Run(pCtx, testHook.RunCommand, repoDir).Times(0)
		postWhWorkingDirLocker.EXPECT().TryLock(fixtures.GithubRepo.FullName, newPull.Num, events.DefaultWorkspace).Times(0)
		postWhWorkingDir.EXPECT().Clone(log, fixtures.GithubRepo, newPull, events.DefaultWorkspace).Times(0)
	})
	t.Run("error locking work dir", func(t *testing.T) {
		postWorkflowHooksSetup(t)

		globalCfg := valid.GlobalCfg{
			Repos: []valid.Repo{
				{
					ID: fixtures.GithubRepo.ID(),
					PostWorkflowHooks: []*valid.WorkflowHook{
						&testHook,
					},
				},
			},
		}

		postWh.GlobalCfg = globalCfg

		postWhWorkingDirLocker.EXPECT().TryLock(fixtures.GithubRepo.FullName, newPull.Num, events.DefaultWorkspace, events.DefaultRepoRelDir).Return(func() {}, errors.New("some error"))

		err := postWh.RunPostHooks(ctx, nil)

		Assert(t, err != nil, "error not nil")
		postWhWorkingDir.EXPECT().Clone(log, fixtures.GithubRepo, newPull, events.DefaultWorkspace).Times(0)
		whPostWorkflowHookRunner.EXPECT().Run(pCtx, testHook.RunCommand, repoDir).Times(0)
	})

	t.Run("error cloning", func(t *testing.T) {
		postWorkflowHooksSetup(t)

		unlockCalled := newBool(false)
		unlockFn := func() {
			unlockCalled = newBool(true)
		}

		globalCfg := valid.GlobalCfg{
			Repos: []valid.Repo{
				{
					ID: fixtures.GithubRepo.ID(),
					PostWorkflowHooks: []*valid.WorkflowHook{
						&testHook,
					},
				},
			},
		}

		postWh.GlobalCfg = globalCfg

		postWhWorkingDirLocker.EXPECT().TryLock(fixtures.GithubRepo.FullName, newPull.Num, events.DefaultWorkspace, events.DefaultRepoRelDir).Return(unlockFn, nil)
		postWhWorkingDir.EXPECT().Clone(log, fixtures.GithubRepo, newPull, events.DefaultWorkspace).Return(repoDir, false, errors.New("some error"))

		err := postWh.RunPostHooks(ctx, nil)

		Assert(t, err != nil, "error not nil")

		whPostWorkflowHookRunner.EXPECT().Run(pCtx, testHook.RunCommand, repoDir).Times(0)
		Assert(t, *unlockCalled == true, "unlock function called")
	})

	t.Run("error running post hook", func(t *testing.T) {
		postWorkflowHooksSetup(t)

		unlockCalled := newBool(false)
		unlockFn := func() {
			unlockCalled = newBool(true)
		}

		globalCfg := valid.GlobalCfg{
			Repos: []valid.Repo{
				{
					ID: fixtures.GithubRepo.ID(),
					PostWorkflowHooks: []*valid.WorkflowHook{
						&testHook,
					},
				},
			},
		}

		postWh.GlobalCfg = globalCfg

		postWhWorkingDirLocker.EXPECT().TryLock(fixtures.GithubRepo.FullName, newPull.Num, events.DefaultWorkspace, events.DefaultRepoRelDir).Return(unlockFn, nil)
		postWhWorkingDir.EXPECT().Clone(log, fixtures.GithubRepo, newPull, events.DefaultWorkspace).Return(repoDir, false, nil)
		whPostWorkflowHookRunner.EXPECT().Run(pCtx, testHook.RunCommand, repoDir).Return(result, errors.New("some error"))

		err := postWh.RunPostHooks(ctx, nil)

		Assert(t, err != nil, "error not nil")
		Assert(t, *unlockCalled == true, "unlock function called")
	})

	t.Run("comment args passed to webhooks", func(t *testing.T) {
		postWorkflowHooksSetup(t)

		unlockCalled := newBool(false)
		unlockFn := func() {
			unlockCalled = newBool(true)
		}

		globalCfg := valid.GlobalCfg{
			Repos: []valid.Repo{
				{
					ID: fixtures.GithubRepo.ID(),
					PostWorkflowHooks: []*valid.WorkflowHook{
						&testHook,
					},
				},
			},
		}

		cmd := &events.CommentCommand{
			Flags: []string{"comment", "args"},
		}

		expectedCtx := pCtx
		expectedCtx.EscapedCommentArgs = []string{"\\c\\o\\m\\m\\e\\n\\t", "\\a\\r\\g\\s"}

		postWh.GlobalCfg = globalCfg

		postWhWorkingDirLocker.EXPECT().TryLock(fixtures.GithubRepo.FullName, newPull.Num, events.DefaultWorkspace, events.DefaultRepoRelDir).Return(unlockFn, nil)
		postWhWorkingDir.EXPECT().Clone(log, fixtures.GithubRepo, newPull, events.DefaultWorkspace).Return(repoDir, false, nil)
		whPostWorkflowHookRunner.EXPECT().Run(pCtx, testHook.RunCommand, repoDir).Return(result, nil)

		err := postWh.RunPostHooks(ctx, cmd)

		Ok(t, err)
		whPostWorkflowHookRunner.EXPECT().Run(expectedCtx, testHook.RunCommand, repoDir).Times(1)
		Assert(t, *unlockCalled == true, "unlock function called")
	})
}
