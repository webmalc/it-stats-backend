package db

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/common/test"
)

const databaseKey = "database_uri"

// Should return the config object
func TestNewConfig(t *testing.T) {
	c, _ := NewConfig()
	assert.Contains(t, c.DatabaseURI, "dbname=its_test")
}

// Should return the error
func TestNewConfigError(t *testing.T) {
	o := viper.GetString(databaseKey)
	viper.Set(databaseKey, "")
	defer viper.Set(databaseKey, o)

	_, err := NewConfig()
	assert.Error(t, err)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
