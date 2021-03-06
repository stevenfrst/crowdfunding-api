// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	transaction "github.com/stevenfrst/crowdfunding-api/usecase/transaction"
	mock "github.com/stretchr/testify/mock"
)

// TransactionRepoInterface is an autogenerated mock type for the TransactionRepoInterface type
type TransactionRepoInterface struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: transaksi
func (_m *TransactionRepoInterface) CreateTransaction(transaksi *transaction.Domain) (transaction.Domain, error) {
	ret := _m.Called(transaksi)

	var r0 transaction.Domain
	if rf, ok := ret.Get(0).(func(*transaction.Domain) transaction.Domain); ok {
		r0 = rf(transaksi)
	} else {
		r0 = ret.Get(0).(transaction.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*transaction.Domain) error); ok {
		r1 = rf(transaksi)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ID
func (_m *TransactionRepoInterface) GetByID(ID int) (transaction.Domain, error) {
	ret := _m.Called(ID)

	var r0 transaction.Domain
	if rf, ok := ret.Get(0).(func(int) transaction.Domain); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(transaction.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastTransactionID provides a mock function with given fields:
func (_m *TransactionRepoInterface) GetLastTransactionID() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: _a0
func (_m *TransactionRepoInterface) UpdateTransaction(_a0 *transaction.Domain) (*transaction.Domain, error) {
	ret := _m.Called(_a0)

	var r0 *transaction.Domain
	if rf, ok := ret.Get(0).(func(*transaction.Domain) *transaction.Domain); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*transaction.Domain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
