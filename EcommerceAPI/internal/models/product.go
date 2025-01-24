package models

import "time"

type Product struct {
    ID          string  `gorm:"primaryKey"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
