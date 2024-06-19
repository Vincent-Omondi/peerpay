// config/config.go
package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
