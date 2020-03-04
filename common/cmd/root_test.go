package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-backend/common/mocks"
	"github.com/webmalc/it-stats-backend/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.BaseLogger{}
	cr := NewCommandRouter(m)
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.BaseLogger{}
	cr := NewCommandRouter(m)
	assert.Equal(t, m, cr.logger)
	assert.NotNil(t, cr.rootCmd)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}

// Should add the app commands.
func TestCommandRouter_RegisterApp(t *testing.T) {
	m := &mocks.CommandsAdder{}
	cr := NewCommandRouter(&mocks.BaseLogger{})
	m.On("AddCommands", cr.rootCmd).Return(nil).Once()
	cr.RegisterApp(m)
	m.AssertExpectations(t)
}
