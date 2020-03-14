package mocks

import (
	"github.com/qor/admin"
	"github.com/stretchr/testify/mock"
)

// AdminResource is an autogenerated mock type for the AdminResource type
type AdminResource struct {
	mock.Mock
}

// EditAttrs provides a mock function with given fields: values
func (m *AdminResource) EditAttrs(values ...interface{}) []*admin.Section {
	m.Called(values)
	return []*admin.Section{}
}

// IndexAttrs provides a mock function with given fields: values
func (m *AdminResource) IndexAttrs(values ...interface{}) []*admin.Section {
	m.Called(values)
	return []*admin.Section{}
}

// NewAttrs provides a mock function with given fields: values
func (m *AdminResource) NewAttrs(values ...interface{}) []*admin.Section {
	m.Called(values)
	return []*admin.Section{}
}

// ShowAttrs provides a mock function with given fields: values
func (m *AdminResource) ShowAttrs(values ...interface{}) []*admin.Section {
	m.Called(values)
	return []*admin.Section{}
}

// Filter provides a mock function with given fields: filter
func (m *AdminResource) Filter(filter *admin.Filter) {
	m.Called(filter)
}
