// src/db/user.go
package db

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Age       int
	CreatedAt int64 `gorm:"autoCreateTime"`
}

// AutoMigrate the User table
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
