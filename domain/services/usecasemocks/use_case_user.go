// Code generated by mockery v2.44.1. DO NOT EDIT.

package usecasemocks

import (
	model "arquitectura/domain/model"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UseCaseUser is an autogenerated mock type for the UseCaseUser type
type UseCaseUser struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *UseCaseUser) Create(_a0 model.User) (*uuid.UUID, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(model.User) (*uuid.UUID, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(model.User) *uuid.UUID); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields:
func (_m *UseCaseUser) Get() (model.Users, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 model.Users
	var r1 error
	if rf, ok := ret.Get(0).(func() (model.Users, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() model.Users); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Users)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUseCaseUser creates a new instance of UseCaseUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCaseUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCaseUser {
	mock := &UseCaseUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
