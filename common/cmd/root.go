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

// Runner runs the  API server
type Runner interface {
	Init()
	Run()
}

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger  errorLogger
	api     Runner
	rootCmd *cobra.Command
}

// RegisterApp register the applications.
func (r *CommandRouter) RegisterApp(a app.CommandsAdder) {
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
func NewCommandRouter(log errorLogger, api Runner) CommandRouter {
	return CommandRouter{
		logger:  log,
		api:     api,
		rootCmd: &cobra.Command{Use: "it_stats.app"},
	}
}
