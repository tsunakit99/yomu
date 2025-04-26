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

	articleRepo := infra.NewMarkdownArticleRepository("posts")
	articleUC := usecase.NewArticleUsecase(articleRepo)
	articleHandler := handler.NewArticleHandler(articleUC)

	likeRepo := infra.NewDynamoLikeRepository()
	likeUC := usecase.NewLikeUsecase(likeRepo)
	likeHandler := handler.NewLikeHandler(likeUC)

	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yomu API is healthy 💊")
	})

	e.GET("/api/articles", articleHandler.GetAll)
	e.GET("/api/articles/:slug", articleHandler.GetBySlug)

	e.POST("/api/likes/:slug", likeHandler.AddLike)

	return e
}
