// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"

	query "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, whereClause
func (_m *UserRepository) Find(ctx context.Context, whereClause []query.WhereClause) ([]models.User, error) {
	ret := _m.Called(ctx, whereClause)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, []query.WhereClause) []models.User); ok {
		r0 = rf(ctx, whereClause)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []query.WhereClause) error); ok {
		r1 = rf(ctx, whereClause)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, user
func (_m *UserRepository) Save(ctx context.Context, user *models.User) (*models.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) *models.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
