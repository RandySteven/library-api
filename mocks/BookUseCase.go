// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	query "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	mock "github.com/stretchr/testify/mock"
)

// BookUseCase is an autogenerated mock type for the BookUseCase type
type BookUseCase struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: book
func (_m *BookUseCase) CreateBook(book *models.Book) (*models.Book, error) {
	ret := _m.Called(book)

	var r0 *models.Book
	if rf, ok := ret.Get(0).(func(*models.Book) *models.Book); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBooks provides a mock function with given fields: whereClause
func (_m *BookUseCase) GetAllBooks(whereClause []query.WhereClause) ([]models.Book, error) {
	ret := _m.Called(whereClause)

	var r0 []models.Book
	if rf, ok := ret.Get(0).(func([]query.WhereClause) []models.Book); ok {
		r0 = rf(whereClause)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]query.WhereClause) error); ok {
		r1 = rf(whereClause)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookUseCase creates a new instance of BookUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookUseCase(t mockConstructorTestingTNewBookUseCase) *BookUseCase {
	mock := &BookUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
