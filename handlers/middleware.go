package handlers

import (
	"net/http"

	"avito.ru/shop/middleware"
	"avito.ru/shop/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService     *services.UserService
	PurchaseService *services.PurchaseService
	AuthHandler     *AuthHandler     // Ссылка на обработчик аутентификации
	SendCoinHandler *SendCoinHandler // Ссылка на обработчик перевода монет
	BuyItemHandler  *BuyItemHandler  // Ссылка на обработчик покупки товаров
}

func NewHandler(userService *services.UserService, purchaseService *services.PurchaseService) *Handler {
	return &Handler{
		UserService:     userService,
		PurchaseService: purchaseService,
		AuthHandler:     &AuthHandler{Service: userService},        // Инициализация обработчика аутентификации
		SendCoinHandler: &SendCoinHandler{Service: userService},    // Инициализация обработчика перевода монет
		BuyItemHandler:  &BuyItemHandler{Service: purchaseService}, // Инициализация обработчика покупки товаров
	}
}

func (h *Handler) SetupRoutes(router *gin.Engine, secretKey string) {
	// Hello route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Avito!"})
	})

	// Открытые маршруты
	router.POST("/api/auth", h.AuthHandler.Login) // Используем существующий метод Login

	// Защищенные маршруты
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(secretKey))
	{
		protected.POST("/sendCoin", h.SendCoinHandler.SendCoins) // Используем существующий метод SendCoins
		protected.GET("/buy/:item", h.BuyItemHandler.BuyItem)    // Используем существующий метод BuyItem
	}
}
