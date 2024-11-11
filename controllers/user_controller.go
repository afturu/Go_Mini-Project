package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/entities"
    "tukerin-platform/services"
)

type UserController struct {
    userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
    return &UserController{userService}
}

func (uc *UserController) Register(c echo.Context) error {
    user := new(entities.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := uc.userService.Register(user); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to register user")
    }

    return c.JSON(http.StatusOK, "User registered successfully")
}

func (uc *UserController) Login(c echo.Context) error {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.Bind(&loginData); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    token, err := uc.userService.Login(loginData.Email, loginData.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (uc *UserController) GetUserByID(c echo.Context) error {
    id := c.Param("id")
    user, err := uc.userService.GetUserByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, "User not found")
    }
    return c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
    id := c.Param("id")
    user := new(entities.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := uc.userService.UpdateUser(id, user); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to update user")
    }
    return c.JSON(http.StatusOK, "User updated successfully")
}

func (uc *UserController) DeleteUser(c echo.Context) error {
    id := c.Param("id")
    if err := uc.userService.DeleteUser(id); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to delete user")
    }
    return c.JSON(http.StatusOK, "User deleted successfully")
}