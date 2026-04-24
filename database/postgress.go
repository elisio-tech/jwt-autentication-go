package database

import (
	"fmt"
	"os"

	"github.com/elisio-tech/jwt-autentication-go.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ConnectDB() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "jwt_auth"),
		getEnv("DB_SSLMODE", "disable"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
