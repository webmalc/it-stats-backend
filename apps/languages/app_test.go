package languages

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/internal/db"
	"github.com/webmalc/it-stats-backend/internal/test"
)

func TestApp_Migrate(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	conn.DropTableIfExists(&Language{})
	assert.False(t, conn.HasTable(&Language{}))
	app := NewApp(conn)
	app.Migrate()
	assert.True(t, conn.HasTable(&Language{}))
}

func TestApp_AddCommands(t *testing.T) {
	rootCmd := &cobra.Command{}
	conn := db.NewConnection()
	defer conn.Close()
	app := NewApp(conn)

	assert.False(t, rootCmd.HasSubCommands())
	app.AddCommands(rootCmd)
	assert.True(t, rootCmd.HasSubCommands())
}

func TestNewApp(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	app := NewApp(conn)
	assert.Equal(t, conn, app.db)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
