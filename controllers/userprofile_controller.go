package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/entities"
    "tukerin-platform/services"
)

type UserProfileController struct {
    profileService *services.UserProfileService
}

func NewUserProfileController(profileService *services.UserProfileService) *UserProfileController {
    return &UserProfileController{profileService}
}

func (upc *UserProfileController) CreateProfile(c echo.Context) error {
    profile := new(entities.UserProfile)
    if err := c.Bind(profile); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := upc.profileService.CreateProfile(profile); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create profile")
    }
    return c.JSON(http.StatusCreated, profile)
}

func (upc *UserProfileController) GetProfileByUserID(c echo.Context) error {
    userID := c.Param("user_id")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, "UserID cannot be empty")
    }

    profile, err := upc.profileService.GetProfileByUserID(userID)
    if err != nil {
        return c.JSON(http.StatusNotFound, "Profile not found")
    }
    return c.JSON(http.StatusOK, profile)
}

func (upc *UserProfileController) UpdateProfile(c echo.Context) error {
    userID := c.Param("user_id")
    if userID == "" {
        return c.JSON(http.StatusBadRequest, "User ID cannot be empty")
    }

    profile := new(entities.UserProfile)
    if err := c.Bind(profile); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := upc.profileService.UpdateProfile(userID, profile); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to update profile: " + err.Error())
    }
    return c.JSON(http.StatusOK, "Profile updated successfully")
}

func (upc *UserProfileController) GetAllProfiles(c echo.Context) error {
    profiles, err := upc.profileService.GetAllProfiles()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to get profiles")
    }
    return c.JSON(http.StatusOK, profiles)
}

func (upc *UserProfileController) DeleteProfile(c echo.Context) error {
    userID := c.Param("user_id")
    if err := upc.profileService.DeleteProfile(userID); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to delete profile")
    }
    return c.JSON(http.StatusOK, "Profile deleted successfully")
}