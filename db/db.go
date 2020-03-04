package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // according to the gorm docs
	"github.com/pkg/errors"
)

// Migratable do the database migrations.
type Migratable interface {
	Migrate()
}

// Database is the database connection.
type Database struct {
	*gorm.DB
}

// RegisterApp register the applications.
func (d *Database) RegisterApp(app Migratable) {
	app.Migrate()
}

// NewConnection returns a new database connection.
func NewConnection() *Database {
	config, err := NewConfig()
	if err != nil {
		panic(errors.Wrap(err, "database"))
	}
	db, err := gorm.Open("postgres", config.DatabaseURI)
	if err != nil {
		panic(errors.Wrap(err, "database"))
	}
	return &Database{db}
}