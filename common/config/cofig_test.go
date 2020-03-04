package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// Should return the config filename based on the environment variable.
func TestGetFilename(t *testing.T) {
	assert.Equal(t, "config", getFilename())

	os.Setenv("ITS_ENV", "test")
	assert.Equal(t, "config.test", getFilename())

	os.Setenv("ITS_ENV", "prod")
	assert.Equal(t, "config.prod", getFilename())
}

// Should setup the main configuration.
func TestSetup(t *testing.T) {
	path := viper.GetString("config_file_path")
	assert.Equal(t, "", path)

	os.Setenv("ITS_ENV", "test")
	Setup()
	path = viper.GetString("config_file_path")
	assert.Contains(t, path, "logs/app.test.log")
}

// Should panic with the invalid environment variable
func TestSetupPanic(t *testing.T) {
	os.Setenv("ITS_ENV", "invalid")
	assert.Panics(t, Setup)
}
