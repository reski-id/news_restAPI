// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "portal/domain"

	mock "github.com/stretchr/testify/mock"
)

// ContentData is an autogenerated mock type for the ContentData type
type ContentData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: IDContent
func (_m *ContentData) Delete(IDContent int) bool {
	ret := _m.Called(IDContent)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDContent)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ContentData) GetAll() []domain.Content {
	ret := _m.Called()

	var r0 []domain.Content
	if rf, ok := ret.Get(0).(func() []domain.Content); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Content)
		}
	}

	return r0
}

// GetContentID provides a mock function with given fields: ContentID
func (_m *ContentData) GetContentID(ContentID int) []domain.Content {
	ret := _m.Called(ContentID)

	var r0 []domain.Content
	if rf, ok := ret.Get(0).(func(int) []domain.Content); ok {
		r0 = rf(ContentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Content)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: insertContent
func (_m *ContentData) Insert(insertContent domain.Content) domain.Content {
	ret := _m.Called(insertContent)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(domain.Content) domain.Content); ok {
		r0 = rf(insertContent)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	return r0
}

// Update provides a mock function with given fields: IDContent, updatedContent
func (_m *ContentData) Update(IDContent int, updatedContent domain.Content) domain.Content {
	ret := _m.Called(IDContent, updatedContent)

	var r0 domain.Content
	if rf, ok := ret.Get(0).(func(int, domain.Content) domain.Content); ok {
		r0 = rf(IDContent, updatedContent)
	} else {
		r0 = ret.Get(0).(domain.Content)
	}

	return r0
}

type mockConstructorTestingTNewContentData interface {
	mock.TestingT
	Cleanup(func())
}

// NewContentData creates a new instance of ContentData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContentData(t mockConstructorTestingTNewContentData) *ContentData {
	mock := &ContentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}