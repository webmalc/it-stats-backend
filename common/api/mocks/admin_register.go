package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// AdminRegister is the app register interface
type AdminRegister struct {
	mock.Mock
}

// RegisterApp is method mock
func (m *AdminRegister) RegisterApp(a interface{ ResourcesAdderInterface }) {
	m.Called(a)
}

// MountTo is method mock
func (m *AdminRegister) MountTo(mountTo string, mux *http.ServeMux) {
	m.Called(mountTo, mux)
}
