package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsunakit99/yomu/usecase"
)

type LikeHandler struct {
	uc usecase.LikeUsecase
}

func NewLikeHandler(uc usecase.LikeUsecase) *LikeHandler {
	return &LikeHandler{
		uc: uc,
	}
}

func (h *LikeHandler) AddLike(c echo.Context) error {
	slug := c.Param("slug")
	if slug == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "slug is required"})
	}
	err := h.uc.AddLike(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "like added"})
}
