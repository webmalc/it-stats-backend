package languages

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/apps/languages/mocks"
	"github.com/webmalc/it-stats-backend/common/db"
	"github.com/webmalc/it-stats-backend/common/test"
)

// Should run the migrations.
func TestApp_Migrate(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	conn.DropTableIfExists(&Language{})
	assert.False(t, conn.HasTable(&Language{}))
	app := NewApp(conn)
	app.Migrate()
	assert.True(t, conn.HasTable(&Language{}))
}

// Should add the app commands.
func TestApp_AddCommands(t *testing.T) {
	rootCmd := &cobra.Command{}
	conn := db.NewConnection()
	defer conn.Close()
	app := NewApp(conn)

	assert.False(t, rootCmd.HasSubCommands())
	app.AddCommands(rootCmd)
	assert.True(t, rootCmd.HasSubCommands())
}

// Should add the app admin resources.
func TestApp_AddAdminResources(t *testing.T) {
	adm := &mocks.AdminResourcesRegister{}
	m := &mocks.AdderAdminResources{}
	conn := db.NewConnection()
	defer conn.Close()
	app := NewApp(conn)
	app.admin = adm

	adm.On("Register", m).Return(nil).Once()
	app.AddAdminResources(m)
	adm.AssertExpectations(t)
}

func TestNewApp(t *testing.T) {
	conn := db.NewConnection()
	mixin := &mocks.AdminMixin{}
	defer conn.Close()
	app := NewApp(conn, mixin)
	assert.Equal(t, conn, app.db)
	assert.NotNil(t, app.admin)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
