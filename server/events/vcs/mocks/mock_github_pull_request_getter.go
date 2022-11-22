// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: GithubPullRequestGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v31/github"
	models "github.com/runatlantis/atlantis/server/events/models"
)

// MockGithubPullRequestGetter is a mock of GithubPullRequestGetter interface.
type MockGithubPullRequestGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGithubPullRequestGetterMockRecorder
}

// MockGithubPullRequestGetterMockRecorder is the mock recorder for MockGithubPullRequestGetter.
type MockGithubPullRequestGetterMockRecorder struct {
	mock *MockGithubPullRequestGetter
}

// NewMockGithubPullRequestGetter creates a new mock instance.
func NewMockGithubPullRequestGetter(ctrl *gomock.Controller) *MockGithubPullRequestGetter {
	mock := &MockGithubPullRequestGetter{ctrl: ctrl}
	mock.recorder = &MockGithubPullRequestGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGithubPullRequestGetter) EXPECT() *MockGithubPullRequestGetterMockRecorder {
	return m.recorder
}

// GetPullRequest mocks base method.
func (m *MockGithubPullRequestGetter) GetPullRequest(arg0 models.Repo, arg1 int) (*github.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPullRequest", arg0, arg1)
	ret0, _ := ret[0].(*github.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPullRequest indicates an expected call of GetPullRequest.
func (mr *MockGithubPullRequestGetterMockRecorder) GetPullRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPullRequest", reflect.TypeOf((*MockGithubPullRequestGetter)(nil).GetPullRequest), arg0, arg1)
}
