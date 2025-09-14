package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement"` // Table id
	UserID   string `gorm:"unique"`                   // User ID
	Name     string `gorm:"unique"`                   // User Name
	Email    string `gorm:"unique"`                   // User Email
	Password string `gorm:"size:255"`                 // User Password hash(Argon2)
}
