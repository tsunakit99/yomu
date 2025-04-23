package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsunakit99/yomu/infra"
	"github.com/tsunakit99/yomu/interfaces/handler"
	"github.com/tsunakit99/yomu/usecase"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	articleRepo := infra.NewLocalArticleRepository()
	articleUC := usecase.NewArticleUsecase(articleRepo)
	articleHandler := handler.NewArticleHandler(articleUC)

	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yomu API is healthy 💊")
	})
	e.GET("/api/articles", articleHandler.GetAll)

	return e
}
