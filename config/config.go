package config

import (
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
	dsn := getEnv("DB_URI", "host=localhost user=movie password=movie dbname=movie port=5454 sslmode=disable")

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
