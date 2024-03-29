// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	service "fast-client/service"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// GetLimits provides a mock function with given fields:
func (_m *Service) GetLimits() (uint64, time.Duration) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLimits")
	}

	var r0 uint64
	var r1 time.Duration
	if rf, ok := ret.Get(0).(func() (uint64, time.Duration)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() time.Duration); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(time.Duration)
	}

	return r0, r1
}

// Service_GetLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLimits'
type Service_GetLimits_Call struct {
	*mock.Call
}

// GetLimits is a helper method to define mock.On call
func (_e *Service_Expecter) GetLimits() *Service_GetLimits_Call {
	return &Service_GetLimits_Call{Call: _e.mock.On("GetLimits")}
}

func (_c *Service_GetLimits_Call) Run(run func()) *Service_GetLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Service_GetLimits_Call) Return(n uint64, p time.Duration) *Service_GetLimits_Call {
	_c.Call.Return(n, p)
	return _c
}

func (_c *Service_GetLimits_Call) RunAndReturn(run func() (uint64, time.Duration)) *Service_GetLimits_Call {
	_c.Call.Return(run)
	return _c
}

// Process provides a mock function with given fields: ctx, batch
func (_m *Service) Process(ctx context.Context, batch service.Batch) error {
	ret := _m.Called(ctx, batch)

	if len(ret) == 0 {
		panic("no return value specified for Process")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, service.Batch) error); ok {
		r0 = rf(ctx, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_Process_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Process'
type Service_Process_Call struct {
	*mock.Call
}

// Process is a helper method to define mock.On call
//   - ctx context.Context
//   - batch service.Batch
func (_e *Service_Expecter) Process(ctx interface{}, batch interface{}) *Service_Process_Call {
	return &Service_Process_Call{Call: _e.mock.On("Process", ctx, batch)}
}

func (_c *Service_Process_Call) Run(run func(ctx context.Context, batch service.Batch)) *Service_Process_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(service.Batch))
	})
	return _c
}

func (_c *Service_Process_Call) Return(_a0 error) *Service_Process_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_Process_Call) RunAndReturn(run func(context.Context, service.Batch) error) *Service_Process_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
