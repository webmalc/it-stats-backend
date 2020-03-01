package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/webmalc/it-stats-backend/logger"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger logger.BaseLogger
}

// server runs server.
func (r *CommandRouter) server(cmd *cobra.Command, args []string) {
	fmt.Println("server")
}

// langs gets programming languages.
func (r *CommandRouter) langs(cmd *cobra.Command, args []string) {
	fmt.Println("langs")
}

// Run the router.
func (r *CommandRouter) Run() {
	rootCmd := &cobra.Command{Use: "it_stats.app"}
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run http server",
			Run:   r.server,
		},
	)
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "langs",
			Short: "Get programming languages from the sources",
			Run:   r.langs,
		},
	)
	err := rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter
func NewCommandRouter(log logger.BaseLogger) CommandRouter {
	return CommandRouter{logger: log}
}
