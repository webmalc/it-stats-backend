package cmd

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/webmalc/it-stats-backend/logger"
)

// CommandRouter is the main commands router
type CommandRouter struct {
	logger logger.BaseLogger
}

// Run the router
func (r *CommandRouter) Run() {
	max := 1
	cmdPrint := &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
	For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(max),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	cmdEcho := &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(max),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	err := rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter
func NewCommandRouter(log logger.BaseLogger) CommandRouter {
	return CommandRouter{logger: log}
}
