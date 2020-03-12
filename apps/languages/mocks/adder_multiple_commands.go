package mocks

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

// AdderMultipleCommands is a mock object
type AdderMultipleCommands struct {
	mock.Mock
}

// AddCommand is method mock
func (m *AdderMultipleCommands) AddCommand(args ...*cobra.Command) {
	m.Called(args)
}
