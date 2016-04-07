package server

import (
	"net/http"
	"os"

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

	s.Router.Options("/api/v1/crawl", func(c *echo.Context) error {
		c.Response().Header().Set("Access-Control-Request-Headers", "accept, authorization, content-type")
		c.Response().Header().Set("Access-Control-Request-Method", "POST")
		return c.JSON(http.StatusOK, nil)
	})

	var v1Api = "/api/v1"
	{
		v1ApiGroup := s.Router.Group(v1Api)
		v1ApiGroup.Use(CORSMiddleware())
		v1ApiGroup.Post("/crawl", apiV1Handler.apiV1CrawlUrlPost)
	}

	//s.Router.Run(s.Config.ServerConfig().Address)
	s.Router.Run(":" + os.Getenv("PORT"))
}

func CORSMiddleware() echo.HandlerFunc {
	return func(c *echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, X-Requested-With, Content-Type, Authorization, k-consumer, k-app-type, k-app-version, Total")
		return nil
	}
}
