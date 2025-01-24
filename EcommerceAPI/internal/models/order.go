package models

import "time"

type Order struct{
    ID        string      `gorm:"primaryKey"`
    UserID    string      `gorm:"index"`
    Total     float64     `json:"total"`
    Status    string      `json:"status"`
    Items     []OrderItem `gorm:"foreignKey:OrderID"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type OrderItem struct {
    ID        string  `gorm:"primaryKey"`
    OrderID   string  `gorm:"index"`
    ProductID string  `gorm:"index"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}
