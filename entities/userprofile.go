package entities

type UserProfile struct {
    ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID    uint   `gorm:"not null" json:"user_id"`
    Address   string `gorm:"type:varchar(255)" json:"address"`
    Phone     string `gorm:"type:varchar(20)" json:"phone"`
    Bio       string `gorm:"type:text" json:"bio"`
    ImageURL  string `gorm:"type:varchar(255)" json:"image_url"`
    User      User   `gorm:"foreignKey:UserID"`
}