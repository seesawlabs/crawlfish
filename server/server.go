package server

import (
	"github.com/labstack/echo"
	"github.com/seesawlabs/crawlfish/shared/config"
	"github.com/seesawlabs/crawlfish/shared/firebase"
)

type Server struct {
	Router *echo.Echo
	Config config.IConfiguration
}

func NewServer(c config.IConfiguration) *Server {
	return &Server{
		Config: c.All(),
		Router: echo.New(),
	}
}

func (s *Server) Init() {
	apiV1Handler := &ApiV1CrawlHandler{
		Firebase: firebase.NewFirebase(s.Config.FirebaseConfig()),
	}

	var v1Api = "/api/v1"
	{
		v1ApiGroup := s.Router.Group(v1Api)
		v1ApiGroup.Post("/crawl", apiV1Handler.apiV1CrawlUrlPost)
	}

	s.Router.Run(s.Config.ServerConfig().Address)
}
