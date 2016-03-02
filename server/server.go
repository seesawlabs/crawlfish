package server

import (
	"github.com/labstack/echo"
	"github.com/seesawlabs/crawlfish/shared/config"
	"github.com/seesawlabs/crawlfish/shared/database"
)

type Server struct {
	Router *echo.Echo
	Config config.Server
}

func NewServer(c config.Server) *Server {
	return &Server{
		Config: c,
		Router: echo.New(),
	}
}

func (s *Server) Init(db database.IDatabase) {

	apiV1Handler := &ApiV1CrawlHandler{
		Db: db,
	}

	var v1Api = "/api/v1"
	{
		v1ApiGroup := s.Router.Group(v1Api)
		v1ApiGroup.Get("/crawl/history", apiV1Handler.apiV1CrawlHistoryGet)
		v1ApiGroup.Post("/crawl", apiV1Handler.apiV1CrawlUrlPost)
	}

	s.Router.Run(s.Config.Address)
}
