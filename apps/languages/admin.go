package languages

import (
	"github.com/qor/admin"
	"github.com/webmalc/it-stats-backend/common/app"
)

type adminRegister struct {
	app.Admin
}

// Register registers admin resources.
func (a *adminRegister) Register(adm app.AdderAdminResources) {
	langAdmin := adm.AddResource(&Language{})
	a.RegisterBase(langAdmin)
	sources := GetLanguageSources()

	langAdmin.Meta(&admin.Meta{
		Name:   "Source",
		Config: &admin.SelectOneConfig{Collection: sources},
	})
	langAdmin.Filter(&admin.Filter{
		Name:   "Source",
		Config: &admin.SelectOneConfig{Collection: sources},
	})
}

// newAdmin returns a new admin object.
func newAdmin() *adminRegister {
	adm := &adminRegister{}
	adm.ListFields = []interface{}{
		"ID", "Title", "Position", "Source", "CreatedAt", "UpdatedAt",
	}
	adm.EditFields = []interface{}{"Title", "Position", "Source"}

	return adm
}