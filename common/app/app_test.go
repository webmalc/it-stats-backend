package app

import (
	"testing"

	admin "github.com/qor/admin"
	mock "github.com/stretchr/testify/mock"
)

type mockAdminResource struct {
	mock.Mock
}

func (_m *mockAdminResource) EditAttrs(
	values ...interface{},
) []*admin.Section {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 []*admin.Section
	if rf, ok := ret.Get(0).(func(...interface{}) []*admin.Section); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Section)
		}
	}

	return r0
}

func (_m *mockAdminResource) Filter(filter *admin.Filter) {
	_m.Called(filter)
}

func (_m *mockAdminResource) IndexAttrs(
	values ...interface{},
) []*admin.Section {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 []*admin.Section
	if rf, ok := ret.Get(0).(func(...interface{}) []*admin.Section); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Section)
		}
	}

	return r0
}

func (_m *mockAdminResource) NewAttrs(values ...interface{}) []*admin.Section {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 []*admin.Section
	if rf, ok := ret.Get(0).(func(...interface{}) []*admin.Section); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Section)
		}
	}

	return r0
}

func (_m *mockAdminResource) ShowAttrs(values ...interface{}) []*admin.Section {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 []*admin.Section
	if rf, ok := ret.Get(0).(func(...interface{}) []*admin.Section); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*admin.Section)
		}
	}

	return r0
}

// Should register the base admin attrs and filters.
func TestAdmin_RegisterBase(t *testing.T) {
	m := &mockAdminResource{}
	a := &Admin{
		ListFields: []interface{}{"list one", "list two"},
		EditFields: []interface{}{"edit one", "edit two"},
	}
	m.On("IndexAttrs", a.ListFields...).Return(nil).Once()
	m.On("ShowAttrs", a.ListFields...).Return(nil).Once()
	m.On("NewAttrs", a.EditFields...).Return(nil).Once()
	m.On("EditAttrs", a.EditFields...).Return(nil).Once()
	m.On("Filter", mock.Anything).Return(nil).Twice()
	a.RegisterBase(m)
	m.AssertExpectations(t)
}
