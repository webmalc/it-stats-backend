package languages

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/webmalc/it-stats-backend/common/app"
)

// App is the application structure
type commands struct {
	rootCmd app.AdderMultipleCommands
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
func newCommands(rootCmd app.AdderMultipleCommands) *commands {
	return &commands{rootCmd: rootCmd}
}
