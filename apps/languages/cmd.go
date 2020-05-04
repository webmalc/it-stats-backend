package languages

import (
	"github.com/spf13/cobra"
)

// App is the application structure
type commands struct {
	rootCmd AdderMultipleCommands
	logger  Logger
}

// langs gets programming languages.
func (c *commands) langs(cmd *cobra.Command, args []string) {
	s := NewScrapperRunner(c.logger)
	s.Run()
}

// addCommands adds the app cmd commands.
func (c *commands) addCommands() {
	c.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "langs",
			Short: "Get programming languages from the sources",
			Run:   c.langs,
		},
	)
}

// newCommands returns a new commads object.
func newCommands(logger Logger, rootCmd AdderMultipleCommands) *commands {
	return &commands{logger: logger, rootCmd: rootCmd}
}
