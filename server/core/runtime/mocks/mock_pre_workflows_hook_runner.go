// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: PreWorkflowHookRunner)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

// MockPreWorkflowHookRunner is a mock of PreWorkflowHookRunner interface.
type MockPreWorkflowHookRunner struct {
	ctrl     *gomock.Controller
	recorder *MockPreWorkflowHookRunnerMockRecorder
}

// MockPreWorkflowHookRunnerMockRecorder is the mock recorder for MockPreWorkflowHookRunner.
type MockPreWorkflowHookRunnerMockRecorder struct {
	mock *MockPreWorkflowHookRunner
}

// NewMockPreWorkflowHookRunner creates a new mock instance.
func NewMockPreWorkflowHookRunner(ctrl *gomock.Controller) *MockPreWorkflowHookRunner {
	mock := &MockPreWorkflowHookRunner{ctrl: ctrl}
	mock.recorder = &MockPreWorkflowHookRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPreWorkflowHookRunner) EXPECT() *MockPreWorkflowHookRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockPreWorkflowHookRunner) Run(arg0 models.WorkflowHookCommandContext, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockPreWorkflowHookRunnerMockRecorder) Run(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPreWorkflowHookRunner)(nil).Run), arg0, arg1, arg2)
}
