// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/v8/public/model"
	mock "github.com/stretchr/testify/mock"
)

// PreferenceStore is an autogenerated mock type for the PreferenceStore type
type PreferenceStore struct {
	mock.Mock
}

// CleanupFlagsBatch provides a mock function with given fields: limit
func (_m *PreferenceStore) CleanupFlagsBatch(limit int64) (int64, error) {
	ret := _m.Called(limit)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (int64, error)); ok {
		return rf(limit)
	}
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID, category, name
func (_m *PreferenceStore) Delete(userID string, category string, name string) error {
	ret := _m.Called(userID, category, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(userID, category, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategory provides a mock function with given fields: userID, category
func (_m *PreferenceStore) DeleteCategory(userID string, category string) error {
	ret := _m.Called(userID, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategoryAndName provides a mock function with given fields: category, name
func (_m *PreferenceStore) DeleteCategoryAndName(category string, name string) error {
	ret := _m.Called(category, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(category, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedRows provides a mock function with given fields: limit
func (_m *PreferenceStore) DeleteOrphanedRows(limit int) (int64, error) {
	ret := _m.Called(limit)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int64, error)); ok {
		return rf(limit)
	}
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: userID, category, name
func (_m *PreferenceStore) Get(userID string, category string, name string) (*model.Preference, error) {
	ret := _m.Called(userID, category, name)

	var r0 *model.Preference
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*model.Preference, error)); ok {
		return rf(userID, category, name)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *model.Preference); ok {
		r0 = rf(userID, category, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Preference)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(userID, category, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: userID
func (_m *PreferenceStore) GetAll(userID string) (model.Preferences, error) {
	ret := _m.Called(userID)

	var r0 model.Preferences
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.Preferences, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) model.Preferences); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Preferences)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategory provides a mock function with given fields: userID, category
func (_m *PreferenceStore) GetCategory(userID string, category string) (model.Preferences, error) {
	ret := _m.Called(userID, category)

	var r0 model.Preferences
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (model.Preferences, error)); ok {
		return rf(userID, category)
	}
	if rf, ok := ret.Get(0).(func(string, string) model.Preferences); ok {
		r0 = rf(userID, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Preferences)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryAndName provides a mock function with given fields: category, nane
func (_m *PreferenceStore) GetCategoryAndName(category string, nane string) (model.Preferences, error) {
	ret := _m.Called(category, nane)

	var r0 model.Preferences
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (model.Preferences, error)); ok {
		return rf(category, nane)
	}
	if rf, ok := ret.Get(0).(func(string, string) model.Preferences); ok {
		r0 = rf(category, nane)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Preferences)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(category, nane)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteByUser provides a mock function with given fields: userID
func (_m *PreferenceStore) PermanentDeleteByUser(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: preferences
func (_m *PreferenceStore) Save(preferences model.Preferences) error {
	ret := _m.Called(preferences)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Preferences) error); ok {
		r0 = rf(preferences)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPreferenceStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewPreferenceStore creates a new instance of PreferenceStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPreferenceStore(t mockConstructorTestingTNewPreferenceStore) *PreferenceStore {
	mock := &PreferenceStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
