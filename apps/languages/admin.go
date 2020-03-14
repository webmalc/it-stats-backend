package languages

import (
	"github.com/qor/admin"
)

type adminRegister struct {
	mixins []AdminMixin
}

type config struct {
	listFields []interface{}
	editFields []interface{}
}

func (c *config) GetListFields() []interface{} {
	return c.listFields
}

func (c *config) GetEditFields() []interface{} {
	return c.editFields
}

// Register registers admin resources.
func (a *adminRegister) Register(adm interface{ AdderAdminResources }) {
	langAdmin := adm.AddResource(&Language{})
	a.ApplyMixins(langAdmin, &config{
		listFields: []interface{}{
			"ID", "Title", "Position", "Source", "CreatedAt", "UpdatedAt",
		},
		editFields: []interface{}{"Title", "Position", "Source"},
	})
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

func (a *adminRegister) ApplyMixins(
	resource interface{ AdminResource },
	config AdminConfig,
) {
	for _, mixin := range a.mixins {
		mixin.Apply(resource, config)
	}
}

// newAdmin returns a new admin object.
func newAdmin(mixins ...AdminMixin) *adminRegister {
	return &adminRegister{mixins: mixins}
}
