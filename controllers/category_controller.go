package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/entities"
    "tukerin-platform/services"
)

type CategoryController struct {
    categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
    return &CategoryController{categoryService}
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
    category := new(entities.Category)
    if err := c.Bind(category); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := cc.categoryService.CreateCategory(category); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create category")
    }

    return c.JSON(http.StatusOK, "Category created successfully")
}

func (cc *CategoryController) GetAllCategories(c echo.Context) error {
    categories, err := cc.categoryService.GetAllCategories()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to retrieve categories")
    }
    return c.JSON(http.StatusOK, categories)
}