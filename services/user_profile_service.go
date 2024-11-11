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

func (ups *UserProfileService) GetProfileByUserID(userID string) (*entities.UserProfile, error) {
    return ups.profileRepo.GetProfileByUserID(userID)
}

func (ups *UserProfileService) UpdateProfile(userID string, profile *entities.UserProfile) error {
    return ups.profileRepo.UpdateProfile(userID, profile)
}