package entities

import (
    "time"
)

type Transaction struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    BuyerID     uint      `gorm:"not null" json:"buyer_id"`
    SellerID    uint      `gorm:"not null" json:"seller_id"`
    ItemID      uint      `gorm:"not null" json:"item_id"`
    Status      string    `gorm:"type:varchar(50);not null" json:"status"` // e.g., "pending", "completed", "cancelled"
    Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    Buyer       User      `gorm:"foreignKey:BuyerID"`
    Seller      User      `gorm:"foreignKey:SellerID"`
    Item        Item      `gorm:"foreignKey:ItemID"`
}