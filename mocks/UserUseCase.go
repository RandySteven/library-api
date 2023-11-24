// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	query "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	mock "github.com/stretchr/testify/mock"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserUseCase) CreateUser(user *models.User) (*models.User, error) {
	ret := _m.Called(user)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(*models.User) *models.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields: whereClause
func (_m *UserUseCase) GetAllUsers(whereClause []query.WhereClause) ([]models.User, error) {
	ret := _m.Called(whereClause)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func([]query.WhereClause) []models.User); ok {
		r0 = rf(whereClause)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
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

type mockConstructorTestingTNewUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUseCase(t mockConstructorTestingTNewUserUseCase) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
