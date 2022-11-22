// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/jobs (interfaces: ProjectStatusUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	command "github.com/runatlantis/atlantis/server/events/command"
	models "github.com/runatlantis/atlantis/server/events/models"
)

// MockProjectStatusUpdater is a mock of ProjectStatusUpdater interface.
type MockProjectStatusUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockProjectStatusUpdaterMockRecorder
}

// MockProjectStatusUpdaterMockRecorder is the mock recorder for MockProjectStatusUpdater.
type MockProjectStatusUpdaterMockRecorder struct {
	mock *MockProjectStatusUpdater
}

// NewMockProjectStatusUpdater creates a new mock instance.
func NewMockProjectStatusUpdater(ctrl *gomock.Controller) *MockProjectStatusUpdater {
	mock := &MockProjectStatusUpdater{ctrl: ctrl}
	mock.recorder = &MockProjectStatusUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectStatusUpdater) EXPECT() *MockProjectStatusUpdaterMockRecorder {
	return m.recorder
}

// UpdateProject mocks base method.
func (m *MockProjectStatusUpdater) UpdateProject(arg0 command.ProjectContext, arg1 command.Name, arg2 models.CommitStatus, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectStatusUpdaterMockRecorder) UpdateProject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectStatusUpdater)(nil).UpdateProject), arg0, arg1, arg2, arg3)
}
