package mocks

import (
	"github.com/qor/admin"
	"github.com/stretchr/testify/mock"
)

// AdderAdminResourcesInterface adds resources to the admin
type AdderAdminResourcesInterface interface {
	AddResource(value interface{}, config ...*admin.Config) *admin.Resource
}

// AdminResourcesRegister register admin resources
type AdminResourcesRegister struct {
	mock.Mock
}

// Register is method mock
func (m *AdminResourcesRegister) Register(
	adm interface{ AdderAdminResourcesInterface },
) {
	m.Called(adm)
}
