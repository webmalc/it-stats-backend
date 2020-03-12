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

// Admin is the base admin struct
type Admin struct {
	ListFields []interface{}
	EditFields []interface{}
}

// AdminResource is the admin resource interface.
type AdminResource interface {
	IndexAttrs(values ...interface{}) []*admin.Section
	ShowAttrs(values ...interface{}) []*admin.Section
	NewAttrs(values ...interface{}) []*admin.Section
	EditAttrs(values ...interface{}) []*admin.Section
	Filter(filter *admin.Filter)
}

// RegisterBase register the base admin attrs and filters.
func (a *Admin) RegisterBase(resource AdminResource) {
	resource.IndexAttrs(a.ListFields...)
	resource.ShowAttrs(a.ListFields...)
	resource.NewAttrs(a.EditFields...)
	resource.EditAttrs(a.EditFields...)
	resource.Filter(&admin.Filter{Name: "CreatedAt"})
	resource.Filter(&admin.Filter{Name: "UpdatedAt"})
}
