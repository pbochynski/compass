// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-incubator/compass/components/connector/internal/apperrors"
import certificates "github.com/kyma-incubator/compass/components/connector/internal/certificates"
import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// SignCSR provides a mock function with given fields: encodedCSR, subject
func (_m *Service) SignCSR(encodedCSR []byte, subject certificates.CSRSubject) (certificates.EncodedCertificateChain, apperrors.AppError) {
	ret := _m.Called(encodedCSR, subject)

	var r0 certificates.EncodedCertificateChain
	if rf, ok := ret.Get(0).(func([]byte, certificates.CSRSubject) certificates.EncodedCertificateChain); ok {
		r0 = rf(encodedCSR, subject)
	} else {
		r0 = ret.Get(0).(certificates.EncodedCertificateChain)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func([]byte, certificates.CSRSubject) apperrors.AppError); ok {
		r1 = rf(encodedCSR, subject)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
