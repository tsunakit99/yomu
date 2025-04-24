package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsunakit99/yomu/usecase"
)

type ArticleHandler struct {
	uc usecase.ArticleUsecase
}

func NewArticleHandler(uc usecase.ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{
		uc: uc,
	}
}

func (h *ArticleHandler) GetAll(c echo.Context) error {
	articles, err := h.uc.GetAllArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) GetBySlug(c echo.Context) error {
	slug := c.Param("slug")
	article, err := h.uc.GetArticleBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if article == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Article not found"})
	}
	return c.JSON(http.StatusOK, article)
}
