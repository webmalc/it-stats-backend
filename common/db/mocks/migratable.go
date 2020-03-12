package mocks

import "github.com/stretchr/testify/mock"

// Migratable is the Migratable mock
type Migratable struct {
	mock.Mock
}

// Migrate is a method mock
func (m *Migratable) Migrate() {
	m.Called()
}
