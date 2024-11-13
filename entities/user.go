package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID       uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    Name     string         `gorm:"type:varchar(100);not null" json:"name"`
    Email    string         `gorm:"type:varchar(100);unique;not null" json:"email"`
    Password string         `gorm:"type:varchar(255);not null" json:"password"`
    Profile  UserProfile    `gorm:"foreignKey:UserID" json:"profile"`
    Items    []Item         `gorm:"foreignKey:UserID" json:"items"`
    CreatedAt  *time.Time    `json:"created_at,omitempty"`
    UpdatedAt  *time.Time    `json:"updated_at,omitempty"`
    DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}