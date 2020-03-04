package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/webmalc/it-stats-backend/common/app"
)

// errorLogger logs errors.
type errorLogger interface {
	Error(args ...interface{})
}

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger  errorLogger
	rootCmd *cobra.Command
}

// RegisterApp register the applications.
func (r *CommandRouter) RegisterApp(a app.CommandsAdder) {
	a.AddCommands(r.rootCmd)
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
