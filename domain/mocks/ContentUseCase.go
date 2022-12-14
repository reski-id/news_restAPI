// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "portal/domain"

	mock "github.com/stretchr/testify/mock"
)

// ContentUseCase is an autogenerated mock type for the ContentUseCase type
type ContentUseCase struct {
	mock.Mock
}

// AddContent provides a mock function with given fields: IDUser, useContent
func (_m *ContentUseCase) AddContent(IDUser int, useContent domain.Content) (domain.Content, error) {
	ret := _m.Called(IDUser, useContent)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(int, domain.Content) domain.Content); ok {
		r0 = rf(IDUser, useContent)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Content) error); ok {
		r1 = rf(IDUser, useContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DelContent provides a mock function with given fields: IDContent
func (_m *ContentUseCase) DelContent(IDContent int) (bool, error) {
	ret := _m.Called(IDContent)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDContent)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(IDContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllN provides a mock function with given fields:
func (_m *ContentUseCase) GetAllN() ([]domain.Content, error) {
	ret := _m.Called()

	var r0 []domain.Content
	if rf, ok := ret.Get(0).(func() []domain.Content); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Content)
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

// GetSpecificContent provides a mock function with given fields: ContentID
func (_m *ContentUseCase) GetSpecificContent(ContentID int) ([]domain.Content, error) {
	ret := _m.Called(ContentID)

	var r0 []domain.Content
	if rf, ok := ret.Get(0).(func(int) []domain.Content); ok {
		r0 = rf(ContentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Content)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ContentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpContent provides a mock function with given fields: IDContent, updateData
func (_m *ContentUseCase) UpContent(IDContent int, updateData domain.Content) (domain.Content, error) {
	ret := _m.Called(IDContent, updateData)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(int, domain.Content) domain.Content); ok {
		r0 = rf(IDContent, updateData)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Content) error); ok {
		r1 = rf(IDContent, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContentUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewContentUseCase creates a new instance of ContentUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContentUseCase(t mockConstructorTestingTNewContentUseCase) *ContentUseCase {
	mock := &ContentUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
