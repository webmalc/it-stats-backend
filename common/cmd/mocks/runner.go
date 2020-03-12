package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Runner logs errors.
type Runner struct {
	mock.Mock
}

// Init is method mock
func (m *Runner) Init() {
	m.Called()
}

// Run is method mock
func (m *Runner) Run() {
	m.Called()
}
