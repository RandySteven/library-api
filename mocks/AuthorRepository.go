// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	mock "github.com/stretchr/testify/mock"
)

// AuthorRepository is an autogenerated mock type for the AuthorRepository type
type AuthorRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields:
func (_m *AuthorRepository) Find() ([]models.Author, error) {
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

// FindAuthorByName provides a mock function with given fields: name
func (_m *AuthorRepository) FindAuthorByName(name string) (*models.Author, error) {
	ret := _m.Called(name)

	var r0 *models.Author
	if rf, ok := ret.Get(0).(func(string) *models.Author); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: author
func (_m *AuthorRepository) Save(author *models.Author) (*models.Author, error) {
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

type mockConstructorTestingTNewAuthorRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthorRepository creates a new instance of AuthorRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthorRepository(t mockConstructorTestingTNewAuthorRepository) *AuthorRepository {
	mock := &AuthorRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
