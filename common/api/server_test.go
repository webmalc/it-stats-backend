package api

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/webmalc/it-stats-backend/common/api/mocks"
)

// ServerTestSuite is a test suite to run
type ServerTestSuite struct {
	suite.Suite
	admin *mocks.AdminRegister
}

// SetupTest setups the suite
func (s *ServerTestSuite) SetupTest() {
	s.admin = &mocks.AdminRegister{}
}

// Should return a CORS handler
func (s *ServerTestSuite) TestServer_getCors() {
	server := NewServer(s.admin)
	cors := server.getCors()
	assert.Equal(s.T(), reflect.TypeOf(cors).Kind(), reflect.Func)
}

// Should return a log writer
func (s *ServerTestSuite) TestServer_getWriter() {
	server := NewServer(s.admin)
	w := server.getWriter()
	assert.NotNil(s.T(), w)
}

// Should initialize the server
func (s *ServerTestSuite) TestServer_initServer() {
	server := NewServer(s.admin)
	server.config.IsProd = true
	assert.Nil(s.T(), server.router)
	server.initServer()
	assert.NotNil(s.T(), server.router)
}

// Should initialize the server and admin
func (s *ServerTestSuite) TestServer_Init() {
	server := NewServer(s.admin)
	server.config.IsProd = true
	assert.Nil(s.T(), server.router)
	s.admin.On("MountTo", "/admin", mock.Anything).Return(nil).Once()
	server.Init()
	assert.NotNil(s.T(), server.router)
	s.admin.AssertExpectations(s.T())
}

// Should initialize the admin
func (s *ServerTestSuite) TestServer_initAdmin() {
	a1 := &mocks.ResourcesAdder{}
	a2 := &mocks.ResourcesAdder{}
	server := NewServer(s.admin, a1, a2)
	server.initServer()
	s.admin.On("MountTo", "/admin", mock.Anything).Return(nil).Once()
	s.admin.On("RegisterApp", a1).Return(nil).Once()
	s.admin.On("RegisterApp", a2).Return(nil).Once()
	server.initAdmin()
	s.admin.AssertExpectations(s.T())
}

// Should run the server
func (s *ServerTestSuite) TestServer_Run() {
	// Implement
}

// Should panic
func (s *ServerTestSuite) TestServer_getWriter_panic() {
	o := viper.Get("log_path")
	defer viper.Set("log_path", o)

	viper.Set("log_path", "")
	server := NewServer(s.admin)
	run := func() {
		_ = server.getWriter()
	}
	assert.Panics(s.T(), run)
}

// Should create a new server object
func (s *ServerTestSuite) TestNewServer() {
	a := &mocks.ResourcesAdder{}
	server := NewServer(s.admin, a)
	assert.NotNil(s.T(), server.admin)
	assert.Equal(s.T(), []ResourcesAdder{a}, server.apps)
}

// Runt the test suite
func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
