// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/webhooks (interfaces: SlackClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	webhooks "github.com/runatlantis/atlantis/server/events/webhooks"
)

// MockSlackClient is a mock of SlackClient interface.
type MockSlackClient struct {
	ctrl     *gomock.Controller
	recorder *MockSlackClientMockRecorder
}

// MockSlackClientMockRecorder is the mock recorder for MockSlackClient.
type MockSlackClientMockRecorder struct {
	mock *MockSlackClient
}

// NewMockSlackClient creates a new mock instance.
func NewMockSlackClient(ctrl *gomock.Controller) *MockSlackClient {
	mock := &MockSlackClient{ctrl: ctrl}
	mock.recorder = &MockSlackClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlackClient) EXPECT() *MockSlackClientMockRecorder {
	return m.recorder
}

// AuthTest mocks base method.
func (m *MockSlackClient) AuthTest() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthTest")
	ret0, _ := ret[0].(error)
	return ret0
}

// AuthTest indicates an expected call of AuthTest.
func (mr *MockSlackClientMockRecorder) AuthTest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthTest", reflect.TypeOf((*MockSlackClient)(nil).AuthTest))
}

// PostMessage mocks base method.
func (m *MockSlackClient) PostMessage(arg0 string, arg1 webhooks.ApplyResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMessage indicates an expected call of PostMessage.
func (mr *MockSlackClientMockRecorder) PostMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockSlackClient)(nil).PostMessage), arg0, arg1)
}

// TokenIsSet mocks base method.
func (m *MockSlackClient) TokenIsSet() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenIsSet")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TokenIsSet indicates an expected call of TokenIsSet.
func (mr *MockSlackClientMockRecorder) TokenIsSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenIsSet", reflect.TypeOf((*MockSlackClient)(nil).TokenIsSet))
}
