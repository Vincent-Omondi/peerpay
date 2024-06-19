// models/admin.go
package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Role   string `gorm:"size:50;not null"`
	User   User
}
