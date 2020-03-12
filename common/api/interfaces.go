package api

import (
	"net/http"

	"github.com/qor/admin"
)

// AdminRegister is the app register interface
type AdminRegister interface {
	RegisterApp(a interface{ ResourcesAdder })
	MountTo(mountTo string, mux *http.ServeMux)
}

// AdderAdminResources adds resources to the admin
type AdderAdminResources interface {
	AddResource(value interface{}, config ...*admin.Config) *admin.Resource
}

// ResourcesAdder add the app admin resources
type ResourcesAdder interface {
	AddAdminResources(admin interface{ AdderAdminResources })
}
