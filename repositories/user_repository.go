package repositories

import (
    "tukerin-platform/entities"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db}
}

func (ur *UserRepository) Register(user *entities.User) error {
    return ur.db.Create(user).Error
}

func (ur *UserRepository) GetUserByID(id string) (*entities.User, error) {
    var user entities.User
    if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (ur *UserRepository) UpdateUser(id string, user *entities.User) error {
    return ur.db.Model(&entities.User{}).Where("id = ?", id).Updates(user).Error
}

func (ur *UserRepository) DeleteUser(id string) error {
    return ur.db.Delete(&entities.User{}, "id = ?", id).Error
}

func (ur *UserRepository) FindByUsername(username string) (*entities.User, error) {
    var user entities.User
    if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}