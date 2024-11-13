package routes

import (
    "github.com/labstack/echo/v4"
    "tukerin-platform/controllers"
    "tukerin-platform/repositories"
    "tukerin-platform/services"
    "tukerin-platform/config"
)

func InitRoutes(e *echo.Echo) {
    db := config.DB

    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    itemRepo := repositories.NewItemRepository(db)
    itemService := services.NewItemService(itemRepo)
    itemController := controllers.NewItemController(itemService)

    categoryRepo := repositories.NewCategoryRepository(db)
    categoryService := services.NewCategoryService(categoryRepo)
    categoryController := controllers.NewCategoryController(categoryService)

    transactionRepo := repositories.NewTransactionRepository(db)
    transactionService := services.NewTransactionService(transactionRepo)
    transactionController := controllers.NewTransactionController(transactionService)

    // Public routes
    e.POST("/register", userController.Register)
    e.POST("/login", userController.Login)

    // Authenticated routes without middleware
    e.GET("/api/users/:id", userController.GetUserByID)
    e.PUT("/api/users/:id", userController.UpdateUser)
    e.DELETE("/api/users/:id", userController.DeleteUser)

    e.POST("/api/items", itemController.CreateItem)
    e.GET("/api/items/:id", itemController.GetItemByID)
    e.PUT("/api/items/:id", itemController.UpdateItem)
    e.DELETE("/api/items/:id", itemController.DeleteItem)
    e.GET("/api/items", itemController.GetAllItems)

    e.POST("/api/categories", categoryController.CreateCategory)
    e.GET("/api/categories/:id", categoryController.GetCategoryByID)
    e.PUT("/api/categories/:id", categoryController.UpdateCategory)
    e.DELETE("/api/categories/:id", categoryController.DeleteCategory)
    e.GET("/api/categories", categoryController.GetAllCategories)

    e.POST("/api/transactions", transactionController.CreateTransaction)
    e.GET("/api/transactions/:id", transactionController.GetTransactionByID)
    e.GET("/api/transactions", transactionController.GetAllTransactions)
}