package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func JWTSecretKey() string {
	return getEnv("JWT_SECRET_KEY", "secret")
}

func BaseUrl() string {
	baseUrl := getEnv("BASE_URL", "http://localhost:8080")
	return baseUrl
}

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "movie"),
		getEnv("DB_PASSWORD", "movie"),
		getEnv("DB_NAME", "movie"),
		getEnv("DB_PORT", "5454"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
