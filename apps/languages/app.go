package languages

import (
	"github.com/spf13/cobra"
	"github.com/webmalc/it-stats-backend/db"
)

// App is the application structure
type App struct {
	db *db.Database
}

// Migrate does the app migrations
func (a *App) Migrate() {
	a.db.AutoMigrate(&Language{})
}

// AddCommands adds the app cmd commands
func (a *App) AddCommands(rootCmd *cobra.Command) {
	cmd := newCommands(rootCmd)
	cmd.addCommands()
}

// func (a *App) AddAdminPages() {
// 	// TODO: implement
// }
// func (a *App) AddRoutes() {
// 	// TODO: implement
// }

// NewApp returns a new app object
func NewApp(conn *db.Database) *App {
	return &App{db: conn}
}
