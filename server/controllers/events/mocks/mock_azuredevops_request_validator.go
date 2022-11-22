// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/controllers/events (interfaces: AzureDevopsRequestValidator)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAzureDevopsRequestValidator is a mock of AzureDevopsRequestValidator interface.
type MockAzureDevopsRequestValidator struct {
	ctrl     *gomock.Controller
	recorder *MockAzureDevopsRequestValidatorMockRecorder
}

// MockAzureDevopsRequestValidatorMockRecorder is the mock recorder for MockAzureDevopsRequestValidator.
type MockAzureDevopsRequestValidatorMockRecorder struct {
	mock *MockAzureDevopsRequestValidator
}

// NewMockAzureDevopsRequestValidator creates a new mock instance.
func NewMockAzureDevopsRequestValidator(ctrl *gomock.Controller) *MockAzureDevopsRequestValidator {
	mock := &MockAzureDevopsRequestValidator{ctrl: ctrl}
	mock.recorder = &MockAzureDevopsRequestValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzureDevopsRequestValidator) EXPECT() *MockAzureDevopsRequestValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockAzureDevopsRequestValidator) Validate(arg0 *http.Request, arg1, arg2 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockAzureDevopsRequestValidatorMockRecorder) Validate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAzureDevopsRequestValidator)(nil).Validate), arg0, arg1, arg2)
}
