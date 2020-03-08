package admin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/common/db"
	"github.com/webmalc/it-stats-backend/common/mocks"
	"github.com/webmalc/it-stats-backend/common/test"
)

// Should register the app resources
func TestAdmin_RegisterApp(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	a := NewAdmin(conn.DB)
	m := &mocks.ResourcesAdder{}
	m.On("AddAdminResources", a.admin).Return(nil).Once()
	a.RegisterApp(m)
	m.AssertExpectations(t)
}

// Should create a new admin object
func TestNewAdmin(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	a := NewAdmin(conn.DB)
	assert.NotNil(t, a.admin)
	assert.Equal(t, "it-stats", a.admin.AdminConfig.SiteName)
}

// Should mounts the admin to the mux
func TestAdmin_MountTo(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	a := NewAdmin(conn.DB)
	assert.NotPanics(t, func() {
		a.MountTo("/admin", http.NewServeMux())
	})
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
