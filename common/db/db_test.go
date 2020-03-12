package db

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/common/db/mocks"
)

// Should create a new database connection.
func TestNewConnection(t *testing.T) {
	conn := NewConnection()
	defer conn.Close()
	assert.NotNil(t, conn)
	err := conn.DB.DB().Ping()
	assert.NoError(t, err)
}

// Should panic.
func TestNewLoggerPanic(t *testing.T) {
	o := viper.Get(databaseKey)
	defer viper.Set(databaseKey, o)

	viper.Set(databaseKey, "")
	assert.Panics(t, func() {
		NewConnection()
	})

	viper.Set(databaseKey, "/invalid/path")
	assert.Panics(t, func() {
		NewConnection()
	})
}

// Should run the app migrations.
func TestDatabase_RegisterApp(t *testing.T) {
	m := &mocks.Migratable{}
	conn := NewConnection()
	defer conn.Close()
	m.On("Migrate").Return(nil).Once()
	conn.RegisterApp(m)
	m.AssertExpectations(t)
}
