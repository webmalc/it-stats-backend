package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // according to the gorm docs
	"github.com/pkg/errors"
)

// Database is the database connection
type Database struct {
	*gorm.DB
}

// NewConnection returns a new database connection
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
