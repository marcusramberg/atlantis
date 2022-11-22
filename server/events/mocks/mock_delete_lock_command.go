// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: DeleteLockCommand)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

// MockDeleteLockCommand is a mock of DeleteLockCommand interface.
type MockDeleteLockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteLockCommandMockRecorder
}

// MockDeleteLockCommandMockRecorder is the mock recorder for MockDeleteLockCommand.
type MockDeleteLockCommandMockRecorder struct {
	mock *MockDeleteLockCommand
}

// NewMockDeleteLockCommand creates a new mock instance.
func NewMockDeleteLockCommand(ctrl *gomock.Controller) *MockDeleteLockCommand {
	mock := &MockDeleteLockCommand{ctrl: ctrl}
	mock.recorder = &MockDeleteLockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteLockCommand) EXPECT() *MockDeleteLockCommandMockRecorder {
	return m.recorder
}

// DeleteLock mocks base method.
func (m *MockDeleteLockCommand) DeleteLock(arg0 string) (*models.ProjectLock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLock", arg0)
	ret0, _ := ret[0].(*models.ProjectLock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLock indicates an expected call of DeleteLock.
func (mr *MockDeleteLockCommandMockRecorder) DeleteLock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLock", reflect.TypeOf((*MockDeleteLockCommand)(nil).DeleteLock), arg0)
}

// DeleteLocksByPull mocks base method.
func (m *MockDeleteLockCommand) DeleteLocksByPull(arg0 string, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLocksByPull", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLocksByPull indicates an expected call of DeleteLocksByPull.
func (mr *MockDeleteLockCommandMockRecorder) DeleteLocksByPull(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLocksByPull", reflect.TypeOf((*MockDeleteLockCommand)(nil).DeleteLocksByPull), arg0, arg1)
}
