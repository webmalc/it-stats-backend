package admin

import "github.com/qor/admin"

// AdderAdminResources adds resources to the admin
type AdderAdminResources interface {
	AddResource(value interface{}, config ...*admin.Config) *admin.Resource
}

// ResourceIndexAdder adds index and show attributes
type ResourceIndexAdder interface {
	IndexAttrs(values ...interface{}) []*admin.Section
	ShowAttrs(values ...interface{}) []*admin.Section
}

// ResourceEditAdder adds new and edit attributes
type ResourceEditAdder interface {
	NewAttrs(values ...interface{}) []*admin.Section
	EditAttrs(values ...interface{}) []*admin.Section
}

// ResourceFilterAdder adds filters
type ResourceFilterAdder interface {
	Filter(filter *admin.Filter)
}

// Resource is the admin resource interface.
type Resource interface {
	ResourceIndexAdder
	ResourceEditAdder
	ResourceFilterAdder
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
