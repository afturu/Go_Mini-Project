package services

import (
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
)

type UserProfileService struct {
    profileRepo repositories.UserProfileRepository
}

func NewUserProfileService(profileRepo repositories.UserProfileRepository) *UserProfileService {
    return &UserProfileService{profileRepo}
}

func (ups *UserProfileService) CreateProfile(profile *entities.UserProfile) error {
    return ups.profileRepo.CreateProfile(profile)
}

func (ups *UserProfileService) GetProfileByUserID(userID string) (*entities.UserProfile, error) {
    return ups.profileRepo.GetProfileByUserID(userID)
}

func (ups *UserProfileService) UpdateProfile(userID string, profile *entities.UserProfile) error {
    return ups.profileRepo.UpdateProfile(userID, profile)
}

func (ups *UserProfileService) GetAllProfiles() ([]entities.UserProfile, error) {
    return ups.profileRepo.GetAllProfiles()
}

func (ups *UserProfileService) DeleteProfile(userID string) error {
    return ups.profileRepo.DeleteProfile(userID)
}