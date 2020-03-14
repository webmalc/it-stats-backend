package admin

import "github.com/qor/admin"

// CrudMixin is a mixin for applying the base CRUD pages
type CrudMixin struct{}

// Apply applyes the mixin to the resource
func (m *CrudMixin) Apply(
	resource interface{ Resource }, config interface{ MixinConfig },
) {
	resource.IndexAttrs(config.GetListFields()...)
	resource.ShowAttrs(config.GetListFields()...)
	resource.NewAttrs(config.GetEditFields()...)
	resource.EditAttrs(config.GetEditFields()...)
	resource.Filter(&admin.Filter{Name: "CreatedAt"})
	resource.Filter(&admin.Filter{Name: "UpdatedAt"})
}
