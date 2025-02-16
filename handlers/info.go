package handlers

import (
	"net/http"

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

	// Получаем информацию о пользователе
	info, err := h.Service.GetUserInfo(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Errors: err.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, info)
}
