package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// The main commands router
type CommandRouter struct {
}

// Runs the router
func (obj *CommandRouter) Run() {
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
		fmt.Println(err)
	}
}

// NewCommandRouter creates a new CommandRouter
func NewCommandRouter() CommandRouter {
	return CommandRouter{}
}
