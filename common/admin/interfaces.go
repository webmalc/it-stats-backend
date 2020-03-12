package admin

import "github.com/qor/admin"

// AdderAdminResources adds resources to the admin
type AdderAdminResources interface {
	AddResource(value interface{}, config ...*admin.Config) *admin.Resource
}

// ResourcesAdder add the app admin resources
type ResourcesAdder interface {
	AddAdminResources(admin interface{ AdderAdminResources })
}
