package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	log.Printf("Parsed user id from context: %v", userID)
	userIDint, err := strconv.ParseInt(userID.(string), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: "Failed to convert user id to int64."})
		return
	}

	user, err := h.Service.UserRepo.GetUserByID(userIDint)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Errors: fmt.Sprintf("User id %v not found", userIDint)})
		return
	}

	err = h.Service.BuyItem(user, itemName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item purchased successfully"})
}
