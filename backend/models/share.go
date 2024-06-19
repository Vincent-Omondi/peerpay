// models/share.go
package models

import (
	"gorm.io/gorm"
)

type Share struct {
	gorm.Model
	UserID uint    `gorm:"not null"`
	Amount float64 `gorm:"not null"`
	Days   int     `gorm:"not null"`
}
