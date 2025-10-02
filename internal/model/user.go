package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    int64     `gorm:"primaryKey;unique"` // User ID
	Name      string    `gorm:"size:32;not null"`  // User Name
	Email     string    `gorm:"size:64;not null"`  // User Email
	Password  string    `gorm:"size:255;not null"` // User Password hash(Argon2)
	CreatedAt time.Time `gorm:"->"`                // Create Time

	_ struct{} `gorm:"uniqueIndex:idx_name_email"`
}
