package languages

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/spf13/cobra"
)

type autoMigrater interface {
	AutoMigrate(values ...interface{}) *gorm.DB
}

// AdderMultipleCommands adds cmd commands to the other command.
type AdderMultipleCommands interface {
	AddCommand(...*cobra.Command)
}

// AdderAdminResources adds resources to the admin
type AdderAdminResources interface {
	AddResource(value interface{}, config ...*admin.Config) *admin.Resource
}

// AdminResourcesRegister register admin resources
type AdminResourcesRegister interface {
	Register(adm interface{ AdderAdminResources })
}

// AdminConfig is a resource admin configuration
type AdminConfig interface {
	GetListFields() []interface{}
	GetEditFields() []interface{}
}

// AdminResource is the admin resource interface.
type AdminResource interface {
	NewAttrs(values ...interface{}) []*admin.Section
	EditAttrs(values ...interface{}) []*admin.Section
	Filter(filter *admin.Filter)
	IndexAttrs(values ...interface{}) []*admin.Section
	ShowAttrs(values ...interface{}) []*admin.Section
}

// AdminMixin is a mixin for applying additional behavior the resource admin
type AdminMixin interface {
	Apply(interface{ AdminResource }, interface{ AdminConfig })
}
