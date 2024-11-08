package entities

type User struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    Email     string `gorm:"unique;not null"`
    CreatedAt int64
    UpdatedAt int64
}