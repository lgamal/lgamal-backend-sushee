// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// AddNewUser provides a mock function with given fields: _a0
func (_m *UserRepository) AddNewUser(_a0 *entity.User) (*entity.User, error) {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddNewUserSession provides a mock function with given fields: s
func (_m *UserRepository) AddNewUserSession(s *entity.Session) (*entity.Session, error) {
	ret := _m.Called(s)

	var r0 *entity.Session
	if rf, ok := ret.Get(0).(func(*entity.Session) *entity.Session); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Session) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckEmailExistence provides a mock function with given fields: _a0
func (_m *UserRepository) CheckEmailExistence(_a0 string) (string, int, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CheckPhoneExistence provides a mock function with given fields: _a0
func (_m *UserRepository) CheckPhoneExistence(_a0 string) (string, int, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CheckUsernameExistence provides a mock function with given fields: _a0
func (_m *UserRepository) CheckUsernameExistence(_a0 string) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailRole provides a mock function with given fields: roleId
func (_m *UserRepository) GetDetailRole(roleId int) (*entity.Role, error) {
	ret := _m.Called(roleId)

	var r0 *entity.Role
	if rf, ok := ret.Get(0).(func(int) *entity.Role); ok {
		r0 = rf(roleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(roleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmailOrUsername provides a mock function with given fields: _a0
func (_m *UserRepository) GetUserByEmailOrUsername(_a0 string) (*entity.User, error) {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: i
func (_m *UserRepository) GetUserById(i int) (*entity.User, error) {
	ret := _m.Called(i)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(int) *entity.User); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserSessionByRefreshToken provides a mock function with given fields: t
func (_m *UserRepository) GetUserSessionByRefreshToken(t string) (*entity.Session, error) {
	ret := _m.Called(t)

	var r0 *entity.Session
	if rf, ok := ret.Get(0).(func(string) *entity.Session); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserDetailsByUsername provides a mock function with given fields: username, updatePremises
func (_m *UserRepository) UpdateUserDetailsByUsername(username string, updatePremises *entity.User) error {
	ret := _m.Called(username, updatePremises)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *entity.User) error); ok {
		r0 = rf(username, updatePremises)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
