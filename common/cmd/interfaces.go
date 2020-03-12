package cmd

import "github.com/spf13/cobra"

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// AdderMultipleCommands adds cmd commands to the other command.
type AdderMultipleCommands interface {
	AddCommand(...*cobra.Command)
}

// CommandsAdder add the app commands
type CommandsAdder interface {
	AddCommands(cmd interface{ AdderMultipleCommands })
}

// Runner runs the  API server
type Runner interface {
	Init()
	Run()
}
