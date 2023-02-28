// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	chaintype "github.com/smartcontractkit/chainlink/core/services/keystore/chaintype"

	mock "github.com/stretchr/testify/mock"

	ocr2key "github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
)

// OCR2 is an autogenerated mock type for the OCR2 type
type OCR2 struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *OCR2) Add(key ocr2key.KeyBundle) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(ocr2key.KeyBundle) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: _a0
func (_m *OCR2) Create(_a0 chaintype.ChainType) (ocr2key.KeyBundle, error) {
	ret := _m.Called(_a0)

	var r0 ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func(chaintype.ChainType) (ocr2key.KeyBundle, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(chaintype.ChainType) ocr2key.KeyBundle); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func(chaintype.ChainType) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *OCR2) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureKeys provides a mock function with given fields:
func (_m *OCR2) EnsureKeys() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *OCR2) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *OCR2) Get(id string) (ocr2key.KeyBundle, error) {
	ret := _m.Called(id)

	var r0 ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ocr2key.KeyBundle, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) ocr2key.KeyBundle); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *OCR2) GetAll() ([]ocr2key.KeyBundle, error) {
	ret := _m.Called()

	var r0 []ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]ocr2key.KeyBundle, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []ocr2key.KeyBundle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOfType provides a mock function with given fields: _a0
func (_m *OCR2) GetAllOfType(_a0 chaintype.ChainType) ([]ocr2key.KeyBundle, error) {
	ret := _m.Called(_a0)

	var r0 []ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func(chaintype.ChainType) ([]ocr2key.KeyBundle, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(chaintype.ChainType) []ocr2key.KeyBundle); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func(chaintype.ChainType) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: keyJSON, password
func (_m *OCR2) Import(keyJSON []byte, password string) (ocr2key.KeyBundle, error) {
	ret := _m.Called(keyJSON, password)

	var r0 ocr2key.KeyBundle
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string) (ocr2key.KeyBundle, error)); ok {
		return rf(keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func([]byte, string) ocr2key.KeyBundle); ok {
		r0 = rf(keyJSON, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ocr2key.KeyBundle)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOCR2 interface {
	mock.TestingT
	Cleanup(func())
}

// NewOCR2 creates a new instance of OCR2. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOCR2(t mockConstructorTestingTNewOCR2) *OCR2 {
	mock := &OCR2{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
