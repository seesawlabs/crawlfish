package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *ApiV1CrawlHandler) apiV1CrawlHistoryGet(c *echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{"test": "test"})
}
