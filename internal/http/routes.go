package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AllRoutes() *echo.Echo {

	r := echo.New()

	r.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return r
}
