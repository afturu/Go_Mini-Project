package entities

import (
    "time"
)

type Item struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID      uint      `gorm:"not null" json:"user_id"`
    CategoryID  uint      `gorm:"not null" json:"category_id"`
    Title       string    `gorm:"type:varchar(100);not null" json:"title"`
    Description string    `gorm:"type:text" json:"description"`
    Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
    Status      string    `gorm:"type:varchar(50);not null" json:"status"` // e.g., "available", "sold"
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    User        User      `gorm:"foreignKey:UserID"`
    Category    Category  `gorm:"foreignKey:CategoryID"`
}