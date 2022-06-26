// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/model"
)

// CollectionRepo is an autogenerated mock type for the CollectionRepo type
type CollectionRepo struct {
	mock.Mock
}

type CollectionRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *CollectionRepo) EXPECT() *CollectionRepo_Expecter {
	return &CollectionRepo_Expecter{mock: &_m.Mock}
}

// CountSubjectCollections provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate
func (_m *CollectionRepo) CountSubjectCollections(ctx context.Context, userID model.UserID, subjectType uint8, collectionType domain.CollectionType, showPrivate bool) (int64, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, uint8, domain.CollectionType, bool) int64); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, uint8, domain.CollectionType, bool) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_CountSubjectCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountSubjectCollections'
type CollectionRepo_CountSubjectCollections_Call struct {
	*mock.Call
}

// CountSubjectCollections is a helper method to define mock.On call
//  - ctx context.Context
//  - userID model.UserID
//  - subjectType uint8
//  - collectionType model.CollectionType
//  - showPrivate bool
func (_e *CollectionRepo_Expecter) CountSubjectCollections(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}) *CollectionRepo_CountSubjectCollections_Call {
	return &CollectionRepo_CountSubjectCollections_Call{Call: _e.mock.On("CountSubjectCollections", ctx, userID, subjectType, collectionType, showPrivate)}
}

func (_c *CollectionRepo_CountSubjectCollections_Call) Run(run func(ctx context.Context, userID model.UserID, subjectType uint8, collectionType domain.CollectionType, showPrivate bool)) *CollectionRepo_CountSubjectCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(uint8), args[3].(domain.CollectionType), args[4].(bool))
	})
	return _c
}

func (_c *CollectionRepo_CountSubjectCollections_Call) Return(_a0 int64, _a1 error) *CollectionRepo_CountSubjectCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetEpisodeCollection provides a mock function with given fields: ctx, userID, subjectID
func (_m *CollectionRepo) GetEpisodeCollection(ctx context.Context, userID model.UserID, subjectID model.SubjectID) (model.EpisodeCollection, error) {
	ret := _m.Called(ctx, userID, subjectID)

	var r0 model.EpisodeCollection
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, model.SubjectID) model.EpisodeCollection); ok {
		r0 = rf(ctx, userID, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.EpisodeCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, model.SubjectID) error); ok {
		r1 = rf(ctx, userID, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_GetEpisodeCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEpisodeCollection'
type CollectionRepo_GetEpisodeCollection_Call struct {
	*mock.Call
}

// GetEpisodeCollection is a helper method to define mock.On call
//  - ctx context.Context
//  - userID model.UserID
//  - subjectID model.SubjectID
func (_e *CollectionRepo_Expecter) GetEpisodeCollection(ctx interface{}, userID interface{}, subjectID interface{}) *CollectionRepo_GetEpisodeCollection_Call {
	return &CollectionRepo_GetEpisodeCollection_Call{Call: _e.mock.On("GetEpisodeCollection", ctx, userID, subjectID)}
}

func (_c *CollectionRepo_GetEpisodeCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectID model.SubjectID)) *CollectionRepo_GetEpisodeCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(model.SubjectID))
	})
	return _c
}

func (_c *CollectionRepo_GetEpisodeCollection_Call) Return(_a0 model.EpisodeCollection, _a1 error) *CollectionRepo_GetEpisodeCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectCollection provides a mock function with given fields: ctx, userID, subjectID
func (_m *CollectionRepo) GetSubjectCollection(ctx context.Context, userID model.UserID, subjectID model.SubjectID) (model.SubjectCollection, error) {
	ret := _m.Called(ctx, userID, subjectID)

	var r0 model.SubjectCollection
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, model.SubjectID) model.SubjectCollection); ok {
		r0 = rf(ctx, userID, subjectID)
	} else {
		r0 = ret.Get(0).(model.SubjectCollection)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, model.SubjectID) error); ok {
		r1 = rf(ctx, userID, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_GetSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectCollection'
type CollectionRepo_GetSubjectCollection_Call struct {
	*mock.Call
}

// GetSubjectCollection is a helper method to define mock.On call
//  - ctx context.Context
//  - userID model.UserID
//  - subjectID model.SubjectID
func (_e *CollectionRepo_Expecter) GetSubjectCollection(ctx interface{}, userID interface{}, subjectID interface{}) *CollectionRepo_GetSubjectCollection_Call {
	return &CollectionRepo_GetSubjectCollection_Call{Call: _e.mock.On("GetSubjectCollection", ctx, userID, subjectID)}
}

func (_c *CollectionRepo_GetSubjectCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectID model.SubjectID)) *CollectionRepo_GetSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(model.SubjectID))
	})
	return _c
}

func (_c *CollectionRepo_GetSubjectCollection_Call) Return(_a0 model.SubjectCollection, _a1 error) *CollectionRepo_GetSubjectCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListSubjectCollection provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate, limit, offset
func (_m *CollectionRepo) ListSubjectCollection(ctx context.Context, userID model.UserID, subjectType uint8, collectionType domain.CollectionType, showPrivate bool, limit int, offset int) ([]model.SubjectCollection, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)

	var r0 []model.SubjectCollection
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, uint8, domain.CollectionType, bool, int, int) []model.SubjectCollection); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SubjectCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserID, uint8, domain.CollectionType, bool, int, int) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_ListSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubjectCollection'
type CollectionRepo_ListSubjectCollection_Call struct {
	*mock.Call
}

// ListSubjectCollection is a helper method to define mock.On call
//  - ctx context.Context
//  - userID model.UserID
//  - subjectType uint8
//  - collectionType model.CollectionType
//  - showPrivate bool
//  - limit int
//  - offset int
func (_e *CollectionRepo_Expecter) ListSubjectCollection(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}, limit interface{}, offset interface{}) *CollectionRepo_ListSubjectCollection_Call {
	return &CollectionRepo_ListSubjectCollection_Call{Call: _e.mock.On("ListSubjectCollection", ctx, userID, subjectType, collectionType, showPrivate, limit, offset)}
}

func (_c *CollectionRepo_ListSubjectCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectType uint8, collectionType domain.CollectionType, showPrivate bool, limit int, offset int)) *CollectionRepo_ListSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(uint8), args[3].(domain.CollectionType), args[4].(bool), args[5].(int), args[6].(int))
	})
	return _c
}

func (_c *CollectionRepo_ListSubjectCollection_Call) Return(_a0 []model.SubjectCollection, _a1 error) *CollectionRepo_ListSubjectCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateSubjectCollection provides a mock function with given fields: ctx, userID, subjectID, data
func (_m *CollectionRepo) UpdateSubjectCollection(ctx context.Context, userID model.UserID, subjectID model.SubjectID, data model.SubjectCollectionUpdate) error {
	ret := _m.Called(ctx, userID, subjectID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserID, model.SubjectID, model.SubjectCollectionUpdate) error); ok {
		r0 = rf(ctx, userID, subjectID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CollectionRepo_UpdateSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSubjectCollection'
type CollectionRepo_UpdateSubjectCollection_Call struct {
	*mock.Call
}

// UpdateSubjectCollection is a helper method to define mock.On call
//  - ctx context.Context
//  - userID model.UserID
//  - subjectID model.SubjectID
//  - data model.SubjectCollectionUpdate
func (_e *CollectionRepo_Expecter) UpdateSubjectCollection(ctx interface{}, userID interface{}, subjectID interface{}, data interface{}) *CollectionRepo_UpdateSubjectCollection_Call {
	return &CollectionRepo_UpdateSubjectCollection_Call{Call: _e.mock.On("UpdateSubjectCollection", ctx, userID, subjectID, data)}
}

func (_c *CollectionRepo_UpdateSubjectCollection_Call) Run(run func(ctx context.Context, userID model.UserID, subjectID model.SubjectID, data model.SubjectCollectionUpdate)) *CollectionRepo_UpdateSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UserID), args[2].(model.SubjectID), args[3].(model.SubjectCollectionUpdate))
	})
	return _c
}

func (_c *CollectionRepo_UpdateSubjectCollection_Call) Return(_a0 error) *CollectionRepo_UpdateSubjectCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCollectionRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollectionRepo creates a new instance of CollectionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollectionRepo(t mockConstructorTestingTNewCollectionRepo) *CollectionRepo {
	mock := &CollectionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
