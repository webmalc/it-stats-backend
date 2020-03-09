package app

import (
	"github.com/qor/admin"
	"github.com/spf13/cobra"
)

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
	Register(adm AdderAdminResources)
}

// Migratable do the database migrations.
type Migratable interface {
	Migrate()
}

// CommandsAdder add the app commands
type CommandsAdder interface {
	AddCommands(cmd AdderMultipleCommands)
}

// ResourcesAdder add the app admin resources
type ResourcesAdder interface {
	AddAdminResources(admin AdderAdminResources)
}

// App in the app interface
type App interface {
	Migratable
	CommandsAdder
	ResourcesAdder
}

// AdminResource is the admin resource interface.
type AdminResource interface {
	IndexAttrs(values ...interface{}) []*admin.Section
	ShowAttrs(values ...interface{}) []*admin.Section
	NewAttrs(values ...interface{}) []*admin.Section
	EditAttrs(values ...interface{}) []*admin.Section
	Filter(filter *admin.Filter)
}

// Admin is the base admin struct
type Admin struct {
	ListFields []interface{}
	EditFields []interface{}
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
