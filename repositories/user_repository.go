package repositories

import (
	"errors"
	"strconv"
	"tukerin-platform/entities"
	"tukerin-platform/middleware"

	"gorm.io/gorm"
)

type UserRepository interface {
    Register(user *entities.User) error
    Login(email, password string) (*entities.User, error)
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

func (r *userRepository) Register(user *entities.User) error {
    hashedPassword, err := middleware.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
    return r.db.Create(user).Error
}

func (r *userRepository) Login(email, password string) (*entities.User, error) {
    var user entities.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }

    if !middleware.CheckPasswordHash(password, user.Password) {
        return nil, errors.New("invalid email or password")
    }
    return &user, nil
}

func (r *userRepository) GetUserByID(id string) (*entities.User, error) {
    // Validasi format ID
    if _, err := strconv.Atoi(id); err != nil {
        return nil, errors.New("invalid ID format")
    }

    // Cari user berdasarkan ID
    var user entities.User
    if err := r.db.First(&user, "id = ?", id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) UpdateUser(id string, user *entities.User) error {
    return r.db.Model(&entities.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) DeleteUser(id string) error {
    return r.db.Delete(&entities.User{}, id).Error
}