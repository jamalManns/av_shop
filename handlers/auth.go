package handlers

import (
	"log"
	"net/http"
	"time"

	"avito.ru/shop/models"
	"avito.ru/shop/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Обработчик для авторизации.
type AuthHandler struct {
	Service *services.UserService
	Secret  []byte
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Errors: "Invalid request"})
		return
	}

	user, err := h.Service.RegisterUser(req.Username) // Автоматическая регистрация
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Errors: "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	tokenString, err := token.SignedString(h.Secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Errors: "Failed to generate token"})
		return
	}
	log.Printf("New user ID created: %v", user.ID)
	c.JSON(http.StatusOK, models.AuthResponse{Token: tokenString})
}
