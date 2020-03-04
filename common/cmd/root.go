package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// errorLogger logs errors.
type errorLogger interface {
	Error(args ...interface{})
}

// AddCommander adds the app cmd commands to the root command.
type AddCommander interface {
	AddCommands(rootCmd *cobra.Command)
}

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger  errorLogger
	rootCmd *cobra.Command
}

// RegisterApp register the applications.
func (r *CommandRouter) RegisterApp(app AddCommander) {
	app.AddCommands(r.rootCmd)
}

// Run the router.
func (r *CommandRouter) Run() {
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(log errorLogger) CommandRouter {
	return CommandRouter{
		logger:  log,
		rootCmd: &cobra.Command{Use: "it_stats.app"},
	}
}
