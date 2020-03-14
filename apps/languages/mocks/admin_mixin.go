package mocks

import (
	"github.com/qor/admin"
	"github.com/stretchr/testify/mock"
)

// AdminResource is the admin resource interface.
type AdminResource interface {
	IndexAttrs(values ...interface{}) []*admin.Section
	ShowAttrs(values ...interface{}) []*admin.Section
	NewAttrs(values ...interface{}) []*admin.Section
	EditAttrs(values ...interface{}) []*admin.Section
	Filter(filter *admin.Filter)
}

// AdminMixin is a mock object
type AdminMixin struct {
	mock.Mock
}

// Apply is method mock
func (m *AdminMixin) Apply(
	resource interface{ AdminResource },
	config interface{ AdminConfigInterface },
) {
	m.Called(resource, config)
}
