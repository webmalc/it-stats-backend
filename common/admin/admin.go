package admin

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

// Admin is the admin struct.
type Admin struct {
	admin *admin.Admin
}

// RegisterApp register the applications.
func (a *Admin) RegisterApp(apl interface{ ResourcesAdder }) {
	apl.AddAdminResources(a.admin)
}

// MountTo mounts the admin to the mux
func (a *Admin) MountTo(mountTo string, mux *http.ServeMux) {
	a.admin.MountTo("/admin", mux)
}

// dashboardRedirect redirects to the main page
func dashboardRedirect(c *admin.Context) {
	http.Redirect(c.Writer, c.Request, "/admin/languages", http.StatusSeeOther)
}

// NewAdmin returns a new admin object.
func NewAdmin(db *gorm.DB) *Admin {
	c := &admin.AdminConfig{
		SiteName: "it-stats",
		DB:       db,
	}

	a := admin.New(c)
	a.GetRouter().Get("/", dashboardRedirect)
	return &Admin{admin: a}
}
