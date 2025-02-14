package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() *Config {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	return &Config{
		DBHost:     getEnv("DATABASE_HOST", "127.0.0.1"),
		DBPort:     getEnv("DATABASE_PORT", "5432"),
		DBUser:     getEnv("DATABASE_USER", "postgres"),
		DBPassword: getEnv("DATABASE_PASSWORD", "postgres"),
		DBName:     getEnv("DATABASE_NAME", "merchstore"),
		JWTSecret:  getEnv("JWT_SECRET", "your_jwt_secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
