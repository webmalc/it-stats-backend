package api

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qor/log"
)

// Server is the server object
type Server struct {
	config *Config
	router *gin.Engine
	apps   []ResourcesAdder
	admin  AdminRegister
}

// GetCors gets the CORS handler
func (s *Server) getCors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	return cors.New(corsConfig)
}

// getWriter return a log writer
func (s *Server) getWriter() io.Writer {
	file, err := os.OpenFile(
		s.config.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600,
	)
	if err != nil {
		panic(errors.Wrap(err, "server"))
	}
	return io.MultiWriter(file, os.Stdout)
}

// initAdmin initializes the the admin
func (s *Server) initAdmin() {
	mux := http.NewServeMux()
	s.admin.MountTo("/admin", mux)
	for _, app := range s.apps {
		s.admin.RegisterApp(app)
	}
	s.router.Any("/admin/*resources", gin.WrapH(mux))
}

// init initializes the router
func (s *Server) initServer() {
	if s.config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = s.getWriter()
	s.router = gin.Default()
	s.router.Use(s.getCors())
	s.router.Use(log.Logger(s.config.LogPath, 90))
}

// Init initializes the server and the admin
func (s *Server) Init() {
	s.initServer()
	s.initAdmin()
}

// Run the server
func (s *Server) Run() {
	err := s.router.Run(s.config.GetURL())
	panic(errors.Wrap(err, "server"))
}

// NewServer returns a new admin object.
func NewServer(admin AdminRegister, apps ...ResourcesAdder) *Server {
	return &Server{
		admin:  admin,
		apps:   apps,
		config: NewConfig(),
	}
}
