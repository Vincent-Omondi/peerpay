// database/migrations.go
package database

import (
	"log"

	"peerpay/backend/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Share{}, &models.Admin{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
