package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware проверяет JWT-токен в заголовке Authorization
// SecretKey используется для подписи JWT
func AuthMiddleware(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем заголовок Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}
		log.Printf("Received token: %s", authHeader)

		// Извлекаем токен из заголовка
		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString = strings.Fields(tokenString[1])
		if len(tokenString) != 1 {
			// Только одна строка в массиве.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		// Парсим токен
		token, err := jwt.Parse(tokenString[0], func(token *jwt.Token) (interface{}, error) {
			// Проверяем метод подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("%v", err)})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid or expired token: %v", tokenString[0])})
			c.Abort()
			return
		}

		// Извлекаем userID из токена
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, exists := claims["userID"]
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}

			// Добавляем userID в контекст как строку
			c.Set("userID", strconv.FormatFloat(userID.(float64), 'f', 0, 64)) // Преобразуем float64 в строку
			log.Printf("Extracted userID from token: %v", userID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}
}
