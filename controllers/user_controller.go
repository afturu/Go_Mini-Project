package controllers

import (
	"net/http"
	"strings"
	"tukerin-platform/entities"
	"tukerin-platform/middleware"
	"tukerin-platform/services"

	"github.com/labstack/echo/v4"
)

type UserController struct {
    userService services.UserService
    jwtUtil     middleware.JwtUsers
}

func NewUserController(userService services.UserService) *UserController {
    return &UserController{
        userService: userService,
        jwtUtil:     middleware.JwtUsers{},
    }
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
    // Bind input data ke struct User
    user := new(entities.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    // Panggil Login dari UserService untuk mendapatkan token dan userID
    token, userID, err := uc.userService.Login(user.Email, user.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, err.Error())
    }

    // Return respons dengan token dan userID
    return c.JSON(http.StatusOK, map[string]interface{}{
        "token":   token,
        "user_id": userID,
    })
}

func (uc *UserController) GetUserByID(c echo.Context) error {
    // Ambil token dari header Authorization
    tokenString := c.Request().Header.Get("Authorization")
    if tokenString == "" {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Token is missing",
        })
    }

    // Parsing token untuk mendapatkan klaim
    claims, err := uc.jwtUtil.ParseJWT(strings.TrimPrefix(tokenString, "Bearer "))
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "message": "Unauthorized access",
        })
    }

    // Ambil User berdasarkan ID dari parameter
    user, err := uc.userService.GetUserByID(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]interface{}{
            "message": "User not found",
        })
    }

    // Return user dan klaim JWT
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user":   user,
        "claims": claims,
    })
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