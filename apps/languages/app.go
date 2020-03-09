package languages

import (
	"github.com/jinzhu/gorm"
	"github.com/webmalc/it-stats-backend/common/app"
)

type autoMigrater interface {
	AutoMigrate(values ...interface{}) *gorm.DB
}

// App is the application structure.
type App struct {
	db    autoMigrater
	admin app.AdminResourcesRegister
}

// Migrate does the app migrations.
func (a *App) Migrate() {
	a.db.AutoMigrate(&Language{})
}

// AddCommands adds the app cmd commands.
func (a *App) AddCommands(rootCmd app.AdderMultipleCommands) {
	cmd := newCommands(rootCmd)
	cmd.addCommands()
}

// AddAdminResources adds the app admin resources commands.
func (a *App) AddAdminResources(adm app.AdderAdminResources) {
	a.admin.Register(adm)
}

// NewApp returns a new app object.
func NewApp(conn autoMigrater) *App {
	return &App{db: conn, admin: newAdmin()}
}
