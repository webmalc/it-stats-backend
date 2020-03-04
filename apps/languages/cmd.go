package languages

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AddCommander adds cmd commands to the root command.
type addMultipleCommander interface {
	AddCommand(...*cobra.Command)
}

// App is the application structure
type commands struct {
	rootCmd addMultipleCommander
}

// server runs server.
func (c *commands) server(cmd *cobra.Command, args []string) {
	fmt.Println("server")
}

// langs gets programming languages.
func (c *commands) langs(cmd *cobra.Command, args []string) {
	fmt.Println("langs")
}

// addCommands adds the app cmd commands.
func (c *commands) addCommands() {
	c.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run http server",
			Run:   c.server,
		},
	)
	c.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "langs",
			Short: "Get programming languages from the sources",
			Run:   c.langs,
		},
	)
}

// newCommands returns a new commads object.
func newCommands(rootCmd addMultipleCommander) *commands {
	return &commands{rootCmd: rootCmd}
}
