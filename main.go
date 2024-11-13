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
    itemRepo := repositories.NewItemRepository(db)
    categoryRepo := repositories.NewCategoryRepository(db)
    transactionRepo := repositories.NewTransactionRepository(db)

    userService := services.NewUserService(userRepo)
    itemService := services.NewItemService(itemRepo)
    categoryService := services.NewCategoryService(categoryRepo)
    transactionService := services.NewTransactionService(transactionRepo)

    userController := controllers.NewUserController(userService)
    itemController := controllers.NewItemController(itemService)
    categoryController := controllers.NewCategoryController(categoryService)
    transactionController := controllers.NewTransactionController(transactionService)

    e := echo.New()

    e.POST("/register", userController.Register)
    e.POST("/login", userController.Login)

    api := e.Group("/api")
    api.Use(middleware.JWTMiddleware())

    api.GET("/users/:id", userController.GetUserByID)
    api.PUT("/users/:id", userController.UpdateUser)
    api.DELETE("/users/:id", userController.DeleteUser)

    api.POST("/items", itemController.CreateItem)
    api.GET("/items/:id", itemController.GetItemByID)
    api.PUT("/items/:id", itemController.UpdateItem)
    api.DELETE("/items/:id", itemController.DeleteItem)
    api.GET("/items", itemController.GetAllItems)

    api.POST("/categories", categoryController.CreateCategory)
    api.GET("/categories/:id", categoryController.GetCategoryByID)
    api.PUT("/categories/:id", categoryController.UpdateCategory)
    api.DELETE("/categories/:id", categoryController.DeleteCategory)
    api.GET("/categories", categoryController.GetAllCategories)

    api.POST("/transactions", transactionController.CreateTransaction)
    api.GET("/transactions/:id", transactionController.GetTransactionByID)
    api.GET("/transactions", transactionController.GetAllTransactions)

    e.Logger.Fatal(e.Start(":8080"))
}