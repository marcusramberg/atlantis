// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/webhooks (interfaces: UnderlyingSlackClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	slack "github.com/slack-go/slack"
)

// MockUnderlyingSlackClient is a mock of UnderlyingSlackClient interface.
type MockUnderlyingSlackClient struct {
	ctrl     *gomock.Controller
	recorder *MockUnderlyingSlackClientMockRecorder
}

// MockUnderlyingSlackClientMockRecorder is the mock recorder for MockUnderlyingSlackClient.
type MockUnderlyingSlackClientMockRecorder struct {
	mock *MockUnderlyingSlackClient
}

// NewMockUnderlyingSlackClient creates a new mock instance.
func NewMockUnderlyingSlackClient(ctrl *gomock.Controller) *MockUnderlyingSlackClient {
	mock := &MockUnderlyingSlackClient{ctrl: ctrl}
	mock.recorder = &MockUnderlyingSlackClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnderlyingSlackClient) EXPECT() *MockUnderlyingSlackClientMockRecorder {
	return m.recorder
}

// AuthTest mocks base method.
func (m *MockUnderlyingSlackClient) AuthTest() (*slack.AuthTestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthTest")
	ret0, _ := ret[0].(*slack.AuthTestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthTest indicates an expected call of AuthTest.
func (mr *MockUnderlyingSlackClientMockRecorder) AuthTest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthTest", reflect.TypeOf((*MockUnderlyingSlackClient)(nil).AuthTest))
}

// GetConversations mocks base method.
func (m *MockUnderlyingSlackClient) GetConversations(arg0 *slack.GetConversationsParameters) ([]slack.Channel, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConversations", arg0)
	ret0, _ := ret[0].([]slack.Channel)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetConversations indicates an expected call of GetConversations.
func (mr *MockUnderlyingSlackClientMockRecorder) GetConversations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConversations", reflect.TypeOf((*MockUnderlyingSlackClient)(nil).GetConversations), arg0)
}

// PostMessage mocks base method.
func (m *MockUnderlyingSlackClient) PostMessage(arg0 string, arg1 ...slack.MsgOption) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostMessage", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PostMessage indicates an expected call of PostMessage.
func (mr *MockUnderlyingSlackClientMockRecorder) PostMessage(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockUnderlyingSlackClient)(nil).PostMessage), varargs...)
}
