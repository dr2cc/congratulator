// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package service

import (
	mock "github.com/stretchr/testify/mock"
)

// NewMockTranslator creates a new instance of MockTranslator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTranslator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTranslator {
	mock := &MockTranslator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockTranslator is an autogenerated mock type for the Translator type
type MockTranslator struct {
	mock.Mock
}

type MockTranslator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTranslator) EXPECT() *MockTranslator_Expecter {
	return &MockTranslator_Expecter{mock: &_m.Mock}
}

// RudimentaryTranslator provides a mock function for the type MockTranslator
func (_mock *MockTranslator) RudimentaryTranslator() string {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for RudimentaryTranslator")
	}

	var r0 string
	if returnFunc, ok := ret.Get(0).(func() string); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Get(0).(string)
	}
	return r0
}

// MockTranslator_RudimentaryTranslator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RudimentaryTranslator'
type MockTranslator_RudimentaryTranslator_Call struct {
	*mock.Call
}

// RudimentaryTranslator is a helper method to define mock.On call
func (_e *MockTranslator_Expecter) RudimentaryTranslator() *MockTranslator_RudimentaryTranslator_Call {
	return &MockTranslator_RudimentaryTranslator_Call{Call: _e.mock.On("RudimentaryTranslator")}
}

func (_c *MockTranslator_RudimentaryTranslator_Call) Run(run func()) *MockTranslator_RudimentaryTranslator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTranslator_RudimentaryTranslator_Call) Return(s string) *MockTranslator_RudimentaryTranslator_Call {
	_c.Call.Return(s)
	return _c
}

func (_c *MockTranslator_RudimentaryTranslator_Call) RunAndReturn(run func() string) *MockTranslator_RudimentaryTranslator_Call {
	_c.Call.Return(run)
	return _c
}
