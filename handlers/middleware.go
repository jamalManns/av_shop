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
	InfoHandler     *InfoHandler
	secret          []byte
}

func NewHandler(userService *services.UserService, purchaseService *services.PurchaseService, infoService *services.InfoService, secret []byte) *Handler {
	return &Handler{
		UserService:     userService,
		PurchaseService: purchaseService,
		AuthHandler:     &AuthHandler{Service: userService, Secret: secret}, // Инициализация обработчика аутентификации
		SendCoinHandler: &SendCoinHandler{Service: userService},             // Инициализация обработчика перевода монет
		BuyItemHandler:  &BuyItemHandler{Service: purchaseService},          // Инициализация обработчика покупки товаров
		InfoHandler:     &InfoHandler{Service: infoService},                 // Инициализация обработчика покупки товаров
		secret:          secret,
	}
}

func (h *Handler) SetupRoutes(router *gin.Engine) {
	// Hello route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Avito!"})
	})

	// Открытые маршруты
	router.POST("/api/auth", h.AuthHandler.Login)

	// Защищенные маршруты
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(h.secret))
	{
		protected.POST("/sendCoin", h.SendCoinHandler.SendCoins)
		protected.GET("/buy/:item", h.BuyItemHandler.BuyItem)
		protected.GET("/info", h.InfoHandler.GetUserInfo)
	}
}
