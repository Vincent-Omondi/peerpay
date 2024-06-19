// database/database.go
package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"peerpay/backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Load environment variables from .env file
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Retrieve the database credentials from environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_HOST", ""),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", ""),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection established")
	}

	models.DB = DB // Assign the connection to models.DB
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
