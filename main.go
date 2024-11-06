package main

import (
    "github.com/labstack/echo/v4"
    "tukerin-platform/config"
    "tukerin-platform/controllers"
    "tukerin-platform/middleware"
    "tukerin-platform/repositories"
    "tukerin-platform/services"
)

func main() {
    config.InitDB()
    db := config.DB

    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    e := echo.New()
    e.Use(middleware.JWTMiddleware())

    e.POST("/register", userController.Register)
    e.POST("/login", userController.Login)
    e.GET("/users/:id", userController.GetUserByID)
    e.PUT("/users/:id", userController.UpdateUser)
    e.DELETE("/users/:id", userController.DeleteUser)

    e.Logger.Fatal(e.Start(":8080"))
}