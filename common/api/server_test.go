package api

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/webmalc/it-stats-backend/common/app"
	"github.com/webmalc/it-stats-backend/common/db"
	"github.com/webmalc/it-stats-backend/common/mocks"
)

// ServerTestSuite is a test suite to run
type ServerTestSuite struct {
	suite.Suite
	db *db.Database
}

// SetupTest setups the suite
func (s *ServerTestSuite) SetupTest() {
	s.db = db.NewConnection()
}

// TearDownTest clears the suite
func (s *ServerTestSuite) TearDownTest() {
	s.db.Close()
}

// Should return a CORS handler
func (s *ServerTestSuite) TestServer_getCors() {
	server := NewServer(s.db)
	cors := server.getCors()
	assert.Equal(s.T(), reflect.TypeOf(cors).Kind(), reflect.Func)
}

// Should return a log writer
func (s *ServerTestSuite) TestServer_getWriter() {
	server := NewServer(s.db)
	w := server.getWriter()
	assert.NotNil(s.T(), w)
}

// Should initialize the server
func (s *ServerTestSuite) TestServer_initServer() {
	server := NewServer(s.db)
	server.config.IsProd = true
	assert.Nil(s.T(), server.router)
	assert.Nil(s.T(), server.admin)
	server.initServer()
	assert.NotNil(s.T(), server.router)
	assert.NotNil(s.T(), server.admin)
}

// Should initialize the server and admin
func (s *ServerTestSuite) TestServer_Init() {
	server := NewServer(s.db)
	server.config.IsProd = true
	assert.Nil(s.T(), server.router)
	assert.Nil(s.T(), server.admin)
	server.Init()
	assert.NotNil(s.T(), server.router)
	assert.NotNil(s.T(), server.admin)
}

// Should initialize the admin
func (s *ServerTestSuite) TestServer_initAdmin() {
	a1 := &mocks.ResourcesAdder{}
	a2 := &mocks.ResourcesAdder{}
	adm := &mocks.AdminRegister{}
	server := NewServer(s.db, a1, a2)
	server.initServer()
	server.admin = adm
	adm.On("MountTo", "/admin", mock.Anything).Return(nil).Once()
	adm.On("RegisterApp", a1).Return(nil).Once()
	adm.On("RegisterApp", a2).Return(nil).Once()
	server.initAdmin()
	adm.AssertExpectations(s.T())
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
	server := NewServer(s.db)
	run := func() {
		_ = server.getWriter()
	}
	assert.Panics(s.T(), run)
}

// Should create a new server object
func (s *ServerTestSuite) TestNewServer() {
	a := &mocks.ResourcesAdder{}
	server := NewServer(s.db, a)
	assert.NotNil(s.T(), server.db)
	assert.Equal(s.T(), []app.ResourcesAdder{a}, server.apps)
}

// Runt the test suite
func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
