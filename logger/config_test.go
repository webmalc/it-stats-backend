package logger

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/test"
)

const filePathKey = "config_file_path"

// Should return the config object
func TestNewConfig(t *testing.T) {
	c, _ := NewConfig()
	assert.Equal(t, true, c.IsDebug)
	assert.Contains(t, c.FilePath, "logs/app.test.log")
}

// Should return the error
func TestNewConfigError(t *testing.T) {
	o := viper.GetString(filePathKey)
	viper.Set(filePathKey, "")
	defer viper.Set(filePathKey, o)

	_, err := NewConfig()
	assert.Error(t, err)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
