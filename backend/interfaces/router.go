package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yomu API is healthy 💊")
	})

	return e
}
