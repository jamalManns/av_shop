package handlers

import (
	"net/http"

	"avito.ru/shop/models"
	"avito.ru/shop/services"
	"github.com/gin-gonic/gin"
)

type SendCoinHandler struct {
	Service *services.UserService
}

func (h *SendCoinHandler) SendCoins(c *gin.Context) {
	var req models.SendCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: "Invalid request"})
		return
	}

	fromUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Errors: "Unauthorized"})
		return
	}

	fromUser, err := h.Service.GetUserByID(fromUserID.(int64))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Errors: "User not found"})
		return
	}

	toUser, err := h.Service.UserRepo.GetUserByUsername(req.ToUser)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Errors: "Recipient not found"})
		return
	}

	err = h.Service.TransferCoins(fromUser, toUser, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coins transferred successfully"})
}
