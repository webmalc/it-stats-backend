package mocks

import (
	"github.com/qor/admin"
	"github.com/stretchr/testify/mock"
)

// AdderAdminResources is a mock object
type AdderAdminResources struct {
	mock.Mock
}

// AddResource is method mock
func (m *AdderAdminResources) AddResource(
	value interface{}, config ...*admin.Config,
) *admin.Resource {
	arg := m.Called(value, config)
	return arg.Get(0).(*admin.Resource)
}
