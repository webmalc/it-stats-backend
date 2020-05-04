package languages

// App is the application structure.
type App struct {
	db     autoMigrater
	admin  AdminResourcesRegister
	logger Logger
}

// Migrate does the app migrations.
func (a *App) Migrate() {
	a.db.AutoMigrate(&Language{})
}

// AddCommands adds the app cmd commands.
func (a *App) AddCommands(rootCmd interface{ AdderMultipleCommands }) {
	cmd := newCommands(a.logger, rootCmd)
	cmd.addCommands()
}

// AddAdminResources adds the app admin resources commands.
func (a *App) AddAdminResources(adm interface{ AdderAdminResources }) {
	a.admin.Register(adm)
}

// NewApp returns a new app object.
func NewApp(conn autoMigrater, logger Logger, mixins ...AdminMixin) *App {
	return &App{db: conn, logger: logger, admin: newAdmin(mixins...)}
}
