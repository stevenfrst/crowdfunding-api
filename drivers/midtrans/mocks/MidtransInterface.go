// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	snap "github.com/midtrans/midtrans-go/snap"
)

// MidtransInterface is an autogenerated mock type for the MidtransInterface type
type MidtransInterface struct {
	mock.Mock
}

// GenerateSnapReq provides a mock function with given fields: id, nominal
func (_m *MidtransInterface) GenerateSnapReq(id int, nominal int) *snap.Request {
	ret := _m.Called(id, nominal)

	var r0 *snap.Request
	if rf, ok := ret.Get(0).(func(int, int) *snap.Request); ok {
		r0 = rf(id, nominal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snap.Request)
		}
	}

	return r0
}

// GetLinkResponse provides a mock function with given fields: id, nominal
func (_m *MidtransInterface) GetLinkResponse(id int, nominal int) *snap.Response {
	ret := _m.Called(id, nominal)

	var r0 *snap.Response
	if rf, ok := ret.Get(0).(func(int, int) *snap.Response); ok {
		r0 = rf(id, nominal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snap.Response)
		}
	}

	return r0
}

// SetupGlobalMidtransConfig provides a mock function with given fields:
func (_m *MidtransInterface) SetupGlobalMidtransConfig() {
	_m.Called()
}
