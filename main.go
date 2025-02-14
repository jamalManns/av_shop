package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"avito.ru/shop/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных (пока без логики)
	log.Println("Connecting to database...")
	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Инициализация Gin-роутера
	router := gin.Default()

	// Пока роуты не добавлены
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Avito!"})
	})

	// Запуск сервера
	port := ":8080"
	log.Printf("Server is running on port %s", port)
	log.Fatal(router.Run(port))
}

func connectDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
