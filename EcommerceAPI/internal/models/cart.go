package models

import "time"

type CartItem struct{
    ID        string  `gorm:"primaryKey"`
    UserID    string  `gorm:"index"`
    ProductID string  `gorm:"index"`
    Quantity  int     `json:"quantity"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
