package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Echo")
	})
	return e
}
