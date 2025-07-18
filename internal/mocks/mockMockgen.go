// Code generated by MockGen. DO NOT EDIT.
// Source: person.go
//
// Generated by this command:
//
//	mockgen -source=person.go -destination=mocks/mockMockgen.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTranslator is a mock of Translator interface.
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
	isgomock struct{}
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator.
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance.
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// RudimentaryTranslator mocks base method.
func (m *MockTranslator) RudimentaryTranslator() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RudimentaryTranslator")
	ret0, _ := ret[0].(string)
	return ret0
}

// RudimentaryTranslator indicates an expected call of RudimentaryTranslator.
func (mr *MockTranslatorMockRecorder) RudimentaryTranslator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RudimentaryTranslator", reflect.TypeOf((*MockTranslator)(nil).RudimentaryTranslator))
}
