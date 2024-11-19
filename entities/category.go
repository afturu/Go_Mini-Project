package entities

type Category struct {
    ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    Name     string `gorm:"type:varchar(100);not null" json:"name"`
    ItemsID  uint   `gorm:"not null" json:"user_id"`
    Items    []Item `gorm:"foreignKey:CategoryID" json:"items"`
}