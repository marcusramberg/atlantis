// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: JobMessageSender)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	command "github.com/runatlantis/atlantis/server/events/command"
)

// MockJobMessageSender is a mock of JobMessageSender interface.
type MockJobMessageSender struct {
	ctrl     *gomock.Controller
	recorder *MockJobMessageSenderMockRecorder
}

// MockJobMessageSenderMockRecorder is the mock recorder for MockJobMessageSender.
type MockJobMessageSenderMockRecorder struct {
	mock *MockJobMessageSender
}

// NewMockJobMessageSender creates a new mock instance.
func NewMockJobMessageSender(ctrl *gomock.Controller) *MockJobMessageSender {
	mock := &MockJobMessageSender{ctrl: ctrl}
	mock.recorder = &MockJobMessageSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobMessageSender) EXPECT() *MockJobMessageSenderMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockJobMessageSender) Send(arg0 command.ProjectContext, arg1 string, arg2 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0, arg1, arg2)
}

// Send indicates an expected call of Send.
func (mr *MockJobMessageSenderMockRecorder) Send(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockJobMessageSender)(nil).Send), arg0, arg1, arg2)
}
