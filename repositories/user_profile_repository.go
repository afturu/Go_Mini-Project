package repositories

import (
    "tukerin-platform/entities"
    "gorm.io/gorm"
)

type UserProfileRepository interface {
    GetProfileByUserID(userID string) (*entities.UserProfile, error)
    UpdateProfile(userID string, profile *entities.UserProfile) error
}

type userProfileRepository struct {
    db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
    return &userProfileRepository{db}
}

func (r *userProfileRepository) GetProfileByUserID(userID string) (*entities.UserProfile, error) {
    var profile entities.UserProfile
    if err := r.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
        return nil, err
    }
    return &profile, nil
}

func (r *userProfileRepository) UpdateProfile(userID string, profile *entities.UserProfile) error {
    return r.db.Model(&entities.UserProfile{}).Where("user_id = ?", userID).Updates(profile).Error
}