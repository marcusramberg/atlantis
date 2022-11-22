// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: PostWorkflowHooksCommandRunner)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	events "github.com/runatlantis/atlantis/server/events"
	command "github.com/runatlantis/atlantis/server/events/command"
)

// MockPostWorkflowHooksCommandRunner is a mock of PostWorkflowHooksCommandRunner interface.
type MockPostWorkflowHooksCommandRunner struct {
	ctrl     *gomock.Controller
	recorder *MockPostWorkflowHooksCommandRunnerMockRecorder
}

// MockPostWorkflowHooksCommandRunnerMockRecorder is the mock recorder for MockPostWorkflowHooksCommandRunner.
type MockPostWorkflowHooksCommandRunnerMockRecorder struct {
	mock *MockPostWorkflowHooksCommandRunner
}

// NewMockPostWorkflowHooksCommandRunner creates a new mock instance.
func NewMockPostWorkflowHooksCommandRunner(ctrl *gomock.Controller) *MockPostWorkflowHooksCommandRunner {
	mock := &MockPostWorkflowHooksCommandRunner{ctrl: ctrl}
	mock.recorder = &MockPostWorkflowHooksCommandRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostWorkflowHooksCommandRunner) EXPECT() *MockPostWorkflowHooksCommandRunnerMockRecorder {
	return m.recorder
}

// RunPostHooks mocks base method.
func (m *MockPostWorkflowHooksCommandRunner) RunPostHooks(arg0 *command.Context, arg1 *events.CommentCommand) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunPostHooks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunPostHooks indicates an expected call of RunPostHooks.
func (mr *MockPostWorkflowHooksCommandRunnerMockRecorder) RunPostHooks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPostHooks", reflect.TypeOf((*MockPostWorkflowHooksCommandRunner)(nil).RunPostHooks), arg0, arg1)
}
