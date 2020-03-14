package languages

// App is the application structure.
type App struct {
	db    autoMigrater
	admin AdminResourcesRegister
}

// Migrate does the app migrations.
func (a *App) Migrate() {
	a.db.AutoMigrate(&Language{})
}

// AddCommands adds the app cmd commands.
func (a *App) AddCommands(rootCmd interface{ AdderMultipleCommands }) {
	cmd := newCommands(rootCmd)
	cmd.addCommands()
}

// AddAdminResources adds the app admin resources commands.
func (a *App) AddAdminResources(adm interface{ AdderAdminResources }) {
	a.admin.Register(adm)
}

// NewApp returns a new app object.
func NewApp(conn autoMigrater, mixins ...AdminMixin) *App {
	return &App{db: conn, admin: newAdmin(mixins...)}
}
