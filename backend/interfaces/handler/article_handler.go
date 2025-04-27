package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsunakit99/yomu/domain/model"
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

// POST /api/articles
func (h *ArticleHandler) Create(c echo.Context) error {
	var input model.ArticleInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.uc.CreateArticle(&input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "article created"})
}

// PUT /api/articles/:slug
func (h *ArticleHandler) Update(c echo.Context) error {
	slug := c.Param("slug")
	var input model.ArticleInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.uc.UpdateArticle(slug, &input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "article updated"})
}

// DELETE /api/articles/:slug
func (h *ArticleHandler) Delete(c echo.Context) error {
	slug := c.Param("slug")

	if err := h.uc.DeleteArticle(slug); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "article deleted"})
}
