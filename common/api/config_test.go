package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/common/test"
)

// Should return the API URL
func TestConfig_GetURL(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "localhost:8080", c.GetURL())
}

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.False(t, c.IsProd)
	assert.Equal(t, "localhost", c.Host)
	assert.Equal(t, uint(8080), c.Port)
}

func TestMain(m *testing.M) {
	test.Run(m)
}
