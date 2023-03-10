// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// CartRepository is an autogenerated mock type for the CartRepository type
type CartRepository struct {
	mock.Mock
}

// AddItemToCart provides a mock function with given fields: c
func (_m *CartRepository) AddItemToCart(c *entity.Cart) (*entity.Cart, error) {
	ret := _m.Called(c)

	var r0 *entity.Cart
	if rf, ok := ret.Get(0).(func(*entity.Cart) *entity.Cart); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Cart) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCart provides a mock function with given fields: username
func (_m *CartRepository) DeleteCart(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCartByCartId provides a mock function with given fields: cartId
func (_m *CartRepository) DeleteCartByCartId(cartId int) error {
	ret := _m.Called(cartId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(cartId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCartByCartId provides a mock function with given fields: cartId
func (_m *CartRepository) GetCartByCartId(cartId int) (*entity.Cart, error) {
	ret := _m.Called(cartId)

	var r0 *entity.Cart
	if rf, ok := ret.Get(0).(func(int) *entity.Cart); ok {
		r0 = rf(cartId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(cartId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCartByCartIds provides a mock function with given fields: cartIds
func (_m *CartRepository) GetCartByCartIds(cartIds []int) (*[]entity.Cart, error) {
	ret := _m.Called(cartIds)

	var r0 *[]entity.Cart
	if rf, ok := ret.Get(0).(func([]int) *[]entity.Cart); ok {
		r0 = rf(cartIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(cartIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCartByUsername provides a mock function with given fields: username
func (_m *CartRepository) GetCartByUsername(username string) (*[]entity.Cart, error) {
	ret := _m.Called(username)

	var r0 *[]entity.Cart
	if rf, ok := ret.Get(0).(func(string) *[]entity.Cart); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCartTotalPriceByCartIds provides a mock function with given fields: cartIds
func (_m *CartRepository) GetCartTotalPriceByCartIds(cartIds []int) (float64, error) {
	ret := _m.Called(cartIds)

	var r0 float64
	if rf, ok := ret.Get(0).(func([]int) float64); ok {
		r0 = rf(cartIds)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(cartIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCartByCartId provides a mock function with given fields: cartId, updatePremises
func (_m *CartRepository) UpdateCartByCartId(cartId int, updatePremises *entity.Cart) error {
	ret := _m.Called(cartId, updatePremises)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *entity.Cart) error); ok {
		r0 = rf(cartId, updatePremises)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCartByCartIds provides a mock function with given fields: cartIds, updatePremises
func (_m *CartRepository) UpdateCartByCartIds(cartIds []int, updatePremises *entity.Cart) error {
	ret := _m.Called(cartIds, updatePremises)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int, *entity.Cart) error); ok {
		r0 = rf(cartIds, updatePremises)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCartRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartRepository creates a new instance of CartRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartRepository(t mockConstructorTestingTNewCartRepository) *CartRepository {
	mock := &CartRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
