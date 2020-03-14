package admin

import "github.com/qor/admin"

// CrudMixin is a mixin for applying the base CRUD pages
type CrudMixin struct{}

// Apply applyes the mixin to the resource
func (m *CrudMixin) Apply(
	resource interface{ Resource },
	config interface{ MixinConfig },
) {
	m.setListFields(resource, config.GetListFields())
	m.setEditFields(resource, config.GetEditFields())
	m.setFilters(resource)
}

func (m *CrudMixin) setListFields(
	resource Resource,
	fields []interface{},
) {
	resource.IndexAttrs(fields...)
	resource.ShowAttrs(fields...)
}

func (m *CrudMixin) setEditFields(
	resource Resource,
	fields []interface{},
) {
	resource.NewAttrs(fields...)
	resource.EditAttrs(fields...)
}

func (m *CrudMixin) setFilters(resource Resource) {
	resource.Filter(&admin.Filter{Name: "CreatedAt"})
	resource.Filter(&admin.Filter{Name: "UpdatedAt"})
}
