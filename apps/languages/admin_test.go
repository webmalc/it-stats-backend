package languages

import (
	"testing"

	"github.com/qor/admin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-backend/apps/languages/mocks"
	"github.com/webmalc/it-stats-backend/common/db"
)

// Should register admin resources.
func Test_adminRegister_Register(t *testing.T) {
	mixin := &mocks.AdminMixin{}
	conn := db.NewConnection()
	defer conn.Close()
	qor := admin.New(&admin.AdminConfig{DB: conn.DB})
	m := &mocks.AdderAdminResources{}
	a := newAdmin(mixin)
	resource := qor.AddResource(&Language{})

	m.On("AddResource", &Language{}, mock.Anything).Return(resource).Once()
	mixin.On("Apply", resource, mock.Anything).Return(nil).Once()

	a.Register(m)
	m.AssertExpectations(t)
	mixin.AssertExpectations(t)
}

// Should apply the admin mixins
func Test_adminRegister_ApplyMixins(t *testing.T) {
	one := &mocks.AdminMixin{}
	two := &mocks.AdminMixin{}
	conn := db.NewConnection()
	defer conn.Close()
	qor := admin.New(&admin.AdminConfig{DB: conn.DB})
	config := &mocks.AdminConfig{}
	resource := qor.AddResource(&Language{})

	a := newAdmin(one, two)
	one.On("Apply", resource, config).Return(nil).Once()
	two.On("Apply", resource, config).Return(nil).Once()
	a.ApplyMixins(resource, config)
	one.AssertExpectations(t)
	two.AssertExpectations(t)
}

// Should return a new admin object.
func Test_newAdmin(t *testing.T) {
	mixin := &mocks.AdminMixin{}
	assert.NotNil(t, newAdmin(mixin))
}
