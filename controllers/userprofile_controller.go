package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/entities"
    "tukerin-platform/services"
)

type UserProfileController struct {
    profileService services.UserProfileService
}

func NewUserProfileController(profileService services.UserProfileService) *UserProfileController {
    return &UserProfileController{profileService}
}

func (upc *UserProfileController) GetProfileByUserID(c echo.Context) error {
    userID := c.Param("user_id")
    profile, err := upc.profileService.GetProfileByUserID(userID)
    if err != nil {
        return c.JSON(http.StatusNotFound, "Profile not found")
    }
    return c.JSON(http.StatusOK, profile)
}

func (upc *UserProfileController) UpdateProfile(c echo.Context) error {
    userID := c.Param("user_id")
    profile := new(entities.UserProfile)
    if err := c.Bind(profile); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := upc.profileService.UpdateProfile(userID, profile); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to update profile")
    }
    return c.JSON(http.StatusOK, "Profile updated successfully")
}