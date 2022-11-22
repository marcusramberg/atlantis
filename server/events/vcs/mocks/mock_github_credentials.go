// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: GithubCredentials)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGithubCredentials is a mock of GithubCredentials interface.
type MockGithubCredentials struct {
	ctrl     *gomock.Controller
	recorder *MockGithubCredentialsMockRecorder
}

// MockGithubCredentialsMockRecorder is the mock recorder for MockGithubCredentials.
type MockGithubCredentialsMockRecorder struct {
	mock *MockGithubCredentials
}

// NewMockGithubCredentials creates a new mock instance.
func NewMockGithubCredentials(ctrl *gomock.Controller) *MockGithubCredentials {
	mock := &MockGithubCredentials{ctrl: ctrl}
	mock.recorder = &MockGithubCredentialsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGithubCredentials) EXPECT() *MockGithubCredentialsMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockGithubCredentials) Client() (*http.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*http.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Client indicates an expected call of Client.
func (mr *MockGithubCredentialsMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockGithubCredentials)(nil).Client))
}

// GetToken mocks base method.
func (m *MockGithubCredentials) GetToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockGithubCredentialsMockRecorder) GetToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockGithubCredentials)(nil).GetToken))
}

// GetUser mocks base method.
func (m *MockGithubCredentials) GetUser() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockGithubCredentialsMockRecorder) GetUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockGithubCredentials)(nil).GetUser))
}
