package handlers

import (
	"net/http"
	"strconv"

	"avito.ru/shop/models"
	"avito.ru/shop/services"

	"github.com/gin-gonic/gin"
)

type InfoHandler struct {
	Service *services.InfoService
}

func (h *InfoHandler) GetUserInfo(c *gin.Context) {
	// Получаем ID текущего пользователя из контекста
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Errors: "Unauthorized"})
		return
	}

	userIDint, err := strconv.ParseInt(userID.(string), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: "Failed to convert user id to int64."})
		return
	}

	// Получаем информацию о пользователе
	info, err := h.Service.GetUserInfo(userIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Errors: err.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, info)
}
