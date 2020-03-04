package admin

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/webmalc/it-stats-backend/common/app"
)

// Admin is the admin struct.
type Admin struct {
	admin *admin.Admin
}

// RegisterApp register the applications.
func (a *Admin) RegisterApp(apl app.ResourcesAdder) {
	apl.AddAdminResources(a.admin)
}

// NewAdmin returns a new admin object.
func NewAdmin(db *gorm.DB, mux *http.ServeMux) *Admin {
	c := &admin.AdminConfig{
		SiteName: "it-stats",
		DB:       db,
	}
	a := admin.New(c)
	a.MountTo("/", mux)
	return &Admin{admin: a}
}
