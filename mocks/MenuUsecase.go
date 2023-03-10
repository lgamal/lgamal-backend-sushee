// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// MenuUsecase is an autogenerated mock type for the MenuUsecase type
type MenuUsecase struct {
	mock.Mock
}

// AddMenu provides a mock function with given fields: reqBody
func (_m *MenuUsecase) AddMenu(reqBody *entity.MenuAddReqBody) (*entity.Menu, error) {
	ret := _m.Called(reqBody)

	var r0 *entity.Menu
	if rf, ok := ret.Get(0).(func(*entity.MenuAddReqBody) *entity.Menu); ok {
		r0 = rf(reqBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.MenuAddReqBody) error); ok {
		r1 = rf(reqBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMenuByMenuId provides a mock function with given fields: menuId
func (_m *MenuUsecase) DeleteMenuByMenuId(menuId int) error {
	ret := _m.Called(menuId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(menuId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMenu provides a mock function with given fields: _a0
func (_m *MenuUsecase) GetMenu(_a0 entity.MenuQuery) (*[]entity.Menu, int, error) {
	ret := _m.Called(_a0)

	var r0 *[]entity.Menu
	if rf, ok := ret.Get(0).(func(entity.MenuQuery) *[]entity.Menu); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Menu)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(entity.MenuQuery) int); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(entity.MenuQuery) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetMenuDetailByMenuId provides a mock function with given fields: menuId
func (_m *MenuUsecase) GetMenuDetailByMenuId(menuId int) (*entity.Menu, error) {
	ret := _m.Called(menuId)

	var r0 *entity.Menu
	if rf, ok := ret.Get(0).(func(int) *entity.Menu); ok {
		r0 = rf(menuId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(menuId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPromotion provides a mock function with given fields:
func (_m *MenuUsecase) GetPromotion() (*[]entity.Promotion, error) {
	ret := _m.Called()

	var r0 *[]entity.Promotion
	if rf, ok := ret.Get(0).(func() *[]entity.Promotion); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Promotion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMenuByMenuId provides a mock function with given fields: menuId, m
func (_m *MenuUsecase) UpdateMenuByMenuId(menuId int, m *entity.Menu) (*entity.Menu, error) {
	ret := _m.Called(menuId, m)

	var r0 *entity.Menu
	if rf, ok := ret.Get(0).(func(int, *entity.Menu) *entity.Menu); ok {
		r0 = rf(menuId, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *entity.Menu) error); ok {
		r1 = rf(menuId, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMenuUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenuUsecase creates a new instance of MenuUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenuUsecase(t mockConstructorTestingTNewMenuUsecase) *MenuUsecase {
	mock := &MenuUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
