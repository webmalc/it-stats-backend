package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-backend/common/cmd/mocks"
	"github.com/webmalc/it-stats-backend/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	cr := NewCommandRouter(m, &mocks.Runner{})
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.ErrorLogger{}
	cr := NewCommandRouter(m, &mocks.Runner{})
	assert.Equal(t, m, cr.logger)
	assert.NotNil(t, cr.rootCmd)
}

// Should add the app commands.
func TestCommandRouter_RegisterApp(t *testing.T) {
	m := &mocks.CommandsAdder{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, &mocks.Runner{})
	m.On("AddCommands", cr.rootCmd).Return(nil).Once()
	cr.RegisterApp(m)
	m.AssertExpectations(t)
}

func TestCommandRouter_server(t *testing.T) {
	m := &mocks.Runner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, m)
	m.On("Run").Return(nil).Once()
	m.On("Init").Return(nil).Once()
	cr.server(&cobra.Command{}, make([]string, 0))
	m.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
