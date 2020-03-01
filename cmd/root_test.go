package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-backend/mocks"
	"github.com/webmalc/it-stats-backend/test"
)

// Should run the root command and log an error
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.BaseLogger{}
	cr := NewCommandRouter(m)
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.BaseLogger{}
	cr := NewCommandRouter(m)
	assert.Equal(t, m, cr.logger)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}

// Should run server
func TestCommandRouter_server(t *testing.T) {
	cr := NewCommandRouter(&mocks.BaseLogger{})
	cr.server(&cobra.Command{}, make([]string, 0))
}

// Should gets languages
func TestCommandRouter_langs(t *testing.T) {
	cr := NewCommandRouter(&mocks.BaseLogger{})
	cr.langs(&cobra.Command{}, make([]string, 0))
}
