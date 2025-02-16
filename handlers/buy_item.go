package handlers

import (
	"net/http"

	"avito.ru/shop/models"
	"avito.ru/shop/services"
	"github.com/gin-gonic/gin"
)

type BuyItemHandler struct {
	Service *services.PurchaseService
}

func (h *BuyItemHandler) BuyItem(c *gin.Context) {
	itemName := c.Param("item")
	if itemName == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: "Item name is required"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Errors: "Unauthorized"})
		return
	}

	user, err := h.Service.UserRepo.GetUserByID(userID.(int64))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Errors: "User not found"})
		return
	}

	err = h.Service.BuyItem(user, itemName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item purchased successfully"})
}
