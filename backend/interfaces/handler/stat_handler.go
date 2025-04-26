package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsunakit99/yomu/usecase"
)

type StatHandler struct {
	uc usecase.StatUsecase
}

func NewStatHandler(uc usecase.StatUsecase) *StatHandler {
	return &StatHandler{uc: uc}
}

func (h *StatHandler) GetStats(c echo.Context) error {
	slug := c.Param("slug")
	if slug == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "slug required"})
	}
	stats, err := h.uc.GetArticleStats(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, stats)
}
