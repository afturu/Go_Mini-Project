package repositories

import (
	"fmt"
	"tukerin-platform/entities"

	"gorm.io/gorm"
)

type UserProfileRepository interface {
    CreateProfile(profile *entities.UserProfile) error 
    GetProfileByUserID(userID string) (*entities.UserProfile, error)
    GetAllProfiles() ([]entities.UserProfile, error)      
    UpdateProfile(userID string, profile *entities.UserProfile) error
    DeleteProfile(userID string) error
}

type userProfileRepository struct {
    db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
    return &userProfileRepository{db}
}

func (r *userProfileRepository) CreateProfile(profile *entities.UserProfile) error {
    return r.db.Create(profile).Error
}

func (r *userProfileRepository) GetProfileByUserID(userID string) (*entities.UserProfile, error) {
    if userID == "" {
        return nil, fmt.Errorf("userID cannot be empty")
    }

    var profile entities.UserProfile
    if err := r.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
        return nil, err
    }
    return &profile, nil
}

func (r *userProfileRepository) GetAllProfiles() ([]entities.UserProfile, error) {
    var profiles []entities.UserProfile
    if err := r.db.Find(&profiles).Error; err != nil {
        return nil, err
    }
    return profiles, nil
}

func (r *userProfileRepository) UpdateProfile(userID string, profile *entities.UserProfile) error {
    if userID == "" {
        return fmt.Errorf("userID cannot be empty")
    }
    return r.db.Model(&entities.UserProfile{}).Where("user_id = ?", userID).Updates(profile).Error
}

func (r *userProfileRepository) DeleteProfile(userID string) error {
    return r.db.Where("user_id = ?", userID).Delete(&entities.UserProfile{}).Error
}

