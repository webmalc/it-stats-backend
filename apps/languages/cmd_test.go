package languages

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-backend/common/mocks"
)

func Test_commands_langs(t *testing.T) {
	m := &mocks.AddMultipleCommander{}
	cmd := newCommands(m)
	cmd.langs(&cobra.Command{}, make([]string, 0))
}

// Should add the app commands to the root command.
func Test_commands_addCommands(t *testing.T) {
	m := &mocks.AddMultipleCommander{}
	cmd := newCommands(m)
	m.On("AddCommand", mock.Anything).Return(nil).Once()
	cmd.addCommands()
	m.AssertExpectations(t)
}

// Should create a new commands object.
func Test_newCommands(t *testing.T) {
	m := &mocks.AddMultipleCommander{}
	cmd := newCommands(m)
	assert.Equal(t, m, cmd.rootCmd)
}
