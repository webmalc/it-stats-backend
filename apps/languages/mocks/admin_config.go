package mocks

import (
	"github.com/stretchr/testify/mock"
)

// AdminConfigInterface is a resource admin configuration
type AdminConfigInterface interface {
	GetListFields() []interface{}
	GetEditFields() []interface{}
}

// AdminConfig is a mock object
type AdminConfig struct {
	mock.Mock
}

// GetListFields is method mock
func (c *AdminConfig) GetListFields() []interface{} {
	arg := c.Called()
	return arg.Get(0).([]interface{})
}

// GetEditFields is method mock
func (c *AdminConfig) GetEditFields() []interface{} {
	arg := c.Called()
	return arg.Get(0).([]interface{})
}
