// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// AuthorHandler is an autogenerated mock type for the AuthorHandler type
type AuthorHandler struct {
	mock.Mock
}

// CreateAuthor provides a mock function with given fields: c
func (_m *AuthorHandler) CreateAuthor(c *gin.Context) {
	_m.Called(c)
}

// GetAllAuthors provides a mock function with given fields: c
func (_m *AuthorHandler) GetAllAuthors(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewAuthorHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthorHandler creates a new instance of AuthorHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthorHandler(t mockConstructorTestingTNewAuthorHandler) *AuthorHandler {
	mock := &AuthorHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
