package mocks

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

// AdderMultipleCommands adds cmd commands to the other command.
type AdderMultipleCommands interface {
	AddCommand(...*cobra.Command)
}

// CommandsAdder is a mock object
type CommandsAdder struct {
	mock.Mock
}

// AddCommands is method mock
func (m *CommandsAdder) AddCommands(cmd interface{ AdderMultipleCommands }) {
	m.Called(cmd)
}
