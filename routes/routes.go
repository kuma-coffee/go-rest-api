package routes

import (
	"net/http"

	"github.com/kuma-coffee/go-rest-api/controllers"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)

	return e
}
