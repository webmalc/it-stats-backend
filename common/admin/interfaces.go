package admin

import "github.com/qor/admin"

// AdderAdminResources adds resources to the admin
type AdderAdminResources interface {
	AddResource(value interface{}, config ...*admin.Config) *admin.Resource
}

// Resource is the admin resource interface.
type Resource interface {
	IndexAttrs(values ...interface{}) []*admin.Section
	ShowAttrs(values ...interface{}) []*admin.Section
	NewAttrs(values ...interface{}) []*admin.Section
	EditAttrs(values ...interface{}) []*admin.Section
	Filter(filter *admin.Filter)
}

// ResourcesAdder add the app admin resources
type ResourcesAdder interface {
	AddAdminResources(admin interface{ AdderAdminResources })
}

// MixinConfig is the admin config for mixins
type MixinConfig interface {
	GetListFields() []interface{}
	GetEditFields() []interface{}
}
