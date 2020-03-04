package test

import (
	"os"
	"testing"

	"github.com/webmalc/it-stats-backend/internal/config"
)

// Setups the tests
func setUp() {
	os.Setenv("ITS_ENV", "test")
	config.Setup()
}

// Run setups, runs and teardowns the tests
func Run(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
