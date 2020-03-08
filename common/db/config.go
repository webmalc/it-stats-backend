package db

import (
	"github.com/spf13/viper"
)

// Config is the database configuration struct.
type Config struct {
	DatabaseURI string
}

// setDefaults sets the default values
func setDefaults() {
	d := "host=localhost port=5432 "
	d += "user=postgres dbname=its password=postgres"
	viper.SetDefault("database_uri", d)
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		DatabaseURI: viper.GetString("database_uri"),
	}
	return config
}
