package main

import (
	"database/sql"
	"fmt"
	"log"

	"avito.ru/shop/config"
	"avito.ru/shop/handlers"
	"avito.ru/shop/repositories"
	"avito.ru/shop/services"
	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// Применение миграций
	err = applyMigrations(cfg)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Инициализация репозиториев
	userRepo := userRepository(db)
	transactionRepo := transactionRepository(db)
	purchaseRepo := purchaseRepository(db)

	// Инициализация сервисов
	userService := services.NewUserService(userRepo, transactionRepo)
	purchaseService := services.NewPurchaseService(userRepo, purchaseRepo)
	infoService := services.NewInfoService(userRepo, transactionRepo, purchaseRepo)

	// Инициализация обработчиков
	handler := handlers.NewHandler(userService, purchaseService, infoService)

	// Инициализация Gin-роутера
	router := gin.Default()

	handler.SetupRoutes(router, cfg.JWTSecret)

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

func applyMigrations(cfg *config.Config) error {
	migrationSource := "file://migrations"
	migrationURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	m, err := migrate.New(migrationSource, migrationURL)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}

func userRepository(db *sql.DB) *repositories.UserRepository {
	return &repositories.UserRepository{DB: db}
}

func transactionRepository(db *sql.DB) *repositories.TransactionRepository {
	return &repositories.TransactionRepository{DB: db}
}

func purchaseRepository(db *sql.DB) *repositories.PurchaseRepository {
	return &repositories.PurchaseRepository{DB: db}
}
