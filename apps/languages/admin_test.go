package languages

import (
	"testing"

	"github.com/qor/admin"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/common/db"
	"github.com/webmalc/it-stats-backend/common/mocks"
)

// Should register admin resources.
func Test_adminRegister_Register(t *testing.T) {
	conn := db.NewConnection()
	qor := admin.New(&admin.AdminConfig{DB: conn.DB})
	m := &mocks.AdderAdminResources{}
	a := newAdmin()

	m.On("AddResource", &Language{}).Return(qor.AddResource(&Language{})).Once()
	a.Register(m)
	m.AssertExpectations(t)
}

// Should return a new admin object.
func Test_newAdmin(t *testing.T) {
	a := newAdmin()
	assert.Contains(t, a.ListFields, "Title")
	assert.Contains(t, a.ListFields, "CreatedAt")
	assert.NotContains(t, a.EditFields, "CreatedAt")
}
