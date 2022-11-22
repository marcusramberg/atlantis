// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: GitlabMergeRequestGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gitlab "github.com/xanzy/go-gitlab"
)

// MockGitlabMergeRequestGetter is a mock of GitlabMergeRequestGetter interface.
type MockGitlabMergeRequestGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGitlabMergeRequestGetterMockRecorder
}

// MockGitlabMergeRequestGetterMockRecorder is the mock recorder for MockGitlabMergeRequestGetter.
type MockGitlabMergeRequestGetterMockRecorder struct {
	mock *MockGitlabMergeRequestGetter
}

// NewMockGitlabMergeRequestGetter creates a new mock instance.
func NewMockGitlabMergeRequestGetter(ctrl *gomock.Controller) *MockGitlabMergeRequestGetter {
	mock := &MockGitlabMergeRequestGetter{ctrl: ctrl}
	mock.recorder = &MockGitlabMergeRequestGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitlabMergeRequestGetter) EXPECT() *MockGitlabMergeRequestGetterMockRecorder {
	return m.recorder
}

// GetMergeRequest mocks base method.
func (m *MockGitlabMergeRequestGetter) GetMergeRequest(arg0 string, arg1 int) (*gitlab.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMergeRequest", arg0, arg1)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMergeRequest indicates an expected call of GetMergeRequest.
func (mr *MockGitlabMergeRequestGetterMockRecorder) GetMergeRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequest", reflect.TypeOf((*MockGitlabMergeRequestGetter)(nil).GetMergeRequest), arg0, arg1)
}
