// models/user.go
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"size:100;not null"`
	Email    string  `gorm:"size:100;unique;not null"`
	Password string  `gorm:"size:100;not null"`
	Profile  Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Shares   []Share
}

type Profile struct {
	gorm.Model
	UserID uint   `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bio    string `gorm:"size:255"`
}

// package models

// import (
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Name     string `gorm:"size:100;not null"`
// 	Email    string `gorm:"size:100;unique;not null"`
// 	Password string `gorm:"size:100;not null"`
// 	Profile  Profile  `gorm:"foreignKey:UserID"`
// 	Shares   []Share
// }

// type Profile struct {
// 	gorm.Model
// 	UserID uint   `gorm:"not null"`
// 	Bio    string `gorm:"size:255"`
// }
