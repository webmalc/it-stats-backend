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
