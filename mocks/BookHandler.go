// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// BookHandler is an autogenerated mock type for the BookHandler type
type BookHandler struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: c
func (_m *BookHandler) CreateBook(c *gin.Context) {
	_m.Called(c)
}

// GetAllBooks provides a mock function with given fields: c
func (_m *BookHandler) GetAllBooks(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewBookHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookHandler creates a new instance of BookHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookHandler(t mockConstructorTestingTNewBookHandler) *BookHandler {
	mock := &BookHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
