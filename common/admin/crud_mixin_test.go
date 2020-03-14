package admin

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-backend/common/admin/mocks"
)

func TestCrudMixin_Apply(t *testing.T) {
	crud := &CrudMixin{}
	resource := &mocks.AdminResource{}
	config := &mocks.MixinConfig{}

	config.On("GetListFields").Twice()
	config.On("GetEditFields").Twice()
	resource.On("IndexAttrs", mock.Anything).Once()
	resource.On("ShowAttrs", mock.Anything).Once()
	resource.On("NewAttrs", mock.Anything).Once()
	resource.On("EditAttrs", mock.Anything).Once()
	resource.On("Filter", mock.Anything).Twice()
	crud.Apply(resource, config)
	config.AssertExpectations(t)
	resource.AssertExpectations(t)
}
