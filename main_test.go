package main

import (
	"os"
	"testing"

	"github.com/webmalc/it-stats-backend/config"
)

// Setups the tests
func setUp() {
	os.Setenv("ITS_ENV", "test")
	config.Setup()
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	// shutdown()
	os.Exit(code)
}
