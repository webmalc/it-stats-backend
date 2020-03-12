package mocks

import (
	"github.com/qor/admin"
	"github.com/stretchr/testify/mock"
)

// AdderAdminResources adds resources to the admin
type AdderAdminResources interface {
	AddResource(value interface{}, config ...*admin.Config) *admin.Resource
}

// ResourcesAdderInterface add the app admin resources
type ResourcesAdderInterface interface {
	AddAdminResources(admin interface{ AdderAdminResources })
}

// ResourcesAdder add the app admin resources
type ResourcesAdder struct {
	mock.Mock
}

// AddAdminResources is a method mock
func (m *ResourcesAdder) AddAdminResources(
	adm interface{ AdderAdminResources },
) {
	m.Called(adm)
}
