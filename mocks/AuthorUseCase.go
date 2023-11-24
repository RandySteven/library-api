// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	mock "github.com/stretchr/testify/mock"
)

// AuthorUseCase is an autogenerated mock type for the AuthorUseCase type
type AuthorUseCase struct {
	mock.Mock
}

// CreateAuthor provides a mock function with given fields: author
func (_m *AuthorUseCase) CreateAuthor(author *models.Author) (*models.Author, error) {
	ret := _m.Called(author)

	var r0 *models.Author
	if rf, ok := ret.Get(0).(func(*models.Author) *models.Author); ok {
		r0 = rf(author)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Author) error); ok {
		r1 = rf(author)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAuthors provides a mock function with given fields:
func (_m *AuthorUseCase) GetAllAuthors() ([]models.Author, error) {
	ret := _m.Called()

	var r0 []models.Author
	if rf, ok := ret.Get(0).(func() []models.Author); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Author)
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

type mockConstructorTestingTNewAuthorUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthorUseCase creates a new instance of AuthorUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthorUseCase(t mockConstructorTestingTNewAuthorUseCase) *AuthorUseCase {
	mock := &AuthorUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}