package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/services"
    "tukerin-platform/entities"
)

type CategoryController struct {
    categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
    return &CategoryController{categoryService}
}

func (c *CategoryController) CreateCategory(ctx echo.Context) error {
    var category entities.Category
    if err := ctx.Bind(&category); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    if err := c.categoryService.CreateCategory(&category); err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) GetCategoryByID(ctx echo.Context) error {
    id := ctx.Param("id")
    category, err := c.categoryService.GetCategoryByID(id)
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) UpdateCategory(ctx echo.Context) error {
    id := ctx.Param("id")
    var category entities.Category
    if err := ctx.Bind(&category); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    if err := c.categoryService.UpdateCategory(id, &category); err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) DeleteCategory(ctx echo.Context) error {
    id := ctx.Param("id")
    if err := c.categoryService.DeleteCategory(id); err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.NoContent(http.StatusOK)
}

func (c *CategoryController) GetAllCategories(ctx echo.Context) error {
    categories, err := c.categoryService.GetAllCategories()
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, categories)
}