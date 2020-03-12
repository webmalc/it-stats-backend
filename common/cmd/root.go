package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger  ErrorLogger
	api     Runner
	rootCmd *cobra.Command
}

// RegisterApp register the applications.
func (r *CommandRouter) RegisterApp(a CommandsAdder) {
	a.AddCommands(r.rootCmd)
}

// server runs server.
func (r *CommandRouter) server(cmd *cobra.Command, args []string) {
	r.api.Init()
	r.api.Run()
}

// Run the router.
func (r *CommandRouter) Run() {
	r.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run the http server",
			Run:   r.server,
		},
	)
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(log ErrorLogger, api Runner) CommandRouter {
	return CommandRouter{
		logger:  log,
		api:     api,
		rootCmd: &cobra.Command{Use: "it_stats.app"},
	}
}
