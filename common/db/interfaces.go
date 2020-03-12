package db

// Migratable do the database migrations.
type Migratable interface {
	Migrate()
}
