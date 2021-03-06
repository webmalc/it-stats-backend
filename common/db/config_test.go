package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/common/test"
)

const databaseKey = "database_uri"

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Contains(t, c.DatabaseURI, "dbname=its_test")
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
