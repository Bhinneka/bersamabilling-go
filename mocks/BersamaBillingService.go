// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	bersamabilling "github.com/Bhinneka/bersamabilling-go"
	mock "github.com/stretchr/testify/mock"
)

// BersamaBillingService is an autogenerated mock type for the BersamaBillingService type
type BersamaBillingService struct {
	mock.Mock
}

// CreatePaymentCode provides a mock function with given fields: request
func (_m *BersamaBillingService) CreatePaymentCode(request bersamabilling.CreatePaymentCodeRequest) (bersamabilling.CreatePaymentCodeResponse, error) {
	ret := _m.Called(request)

	var r0 bersamabilling.CreatePaymentCodeResponse
	if rf, ok := ret.Get(0).(func(bersamabilling.CreatePaymentCodeRequest) bersamabilling.CreatePaymentCodeResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(bersamabilling.CreatePaymentCodeResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bersamabilling.CreatePaymentCodeRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusInquiryPayment provides a mock function with given fields: request
func (_m *BersamaBillingService) StatusInquiryPayment(request bersamabilling.StatusInquiryPaymentRequest) (bersamabilling.StatusInquiryResponse, error) {
	ret := _m.Called(request)

	var r0 bersamabilling.StatusInquiryResponse
	if rf, ok := ret.Get(0).(func(bersamabilling.StatusInquiryPaymentRequest) bersamabilling.StatusInquiryResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(bersamabilling.StatusInquiryResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bersamabilling.StatusInquiryPaymentRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
