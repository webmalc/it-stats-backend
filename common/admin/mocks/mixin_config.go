package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MixinConfig is a mock object
type MixinConfig struct {
	mock.Mock
}

// GetListFields is method mock
func (c *MixinConfig) GetListFields() []interface{} {
	c.Called()
	return []interface{}{}
}

// GetEditFields is method mock
func (c *MixinConfig) GetEditFields() []interface{} {
	c.Called()
	return []interface{}{}
}
