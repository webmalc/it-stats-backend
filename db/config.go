package db

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the database configuration struct.
type Config struct {
	DatabaseURI string
}

// NewConfig returns the configuration object.
func NewConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("database_uri"),
	}
	if config.DatabaseURI == "" {
		return nil, fmt.Errorf("database_uri must be set")
	}
	return config, nil
}
