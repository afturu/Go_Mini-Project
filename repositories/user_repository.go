package repositories

import (
    "tukerin-platform/entities"
    "gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *entities.User) error
    FindByEmail(email string) (*entities.User, error)
    GetUserByID(id string) (*entities.User, error)
    UpdateUser(id string, user *entities.User) error
    DeleteUser(id string) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) CreateUser(user *entities.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
    var user entities.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) GetUserByID(id string) (*entities.User, error) {
    var user entities.User
    if err := r.db.First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) UpdateUser(id string, user *entities.User) error {
    return r.db.Model(&entities.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) DeleteUser(id string) error {
    return r.db.Delete(&entities.User{}, "id = ?", id).Error
}