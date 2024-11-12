package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/entities"
    "tukerin-platform/services"
)

type ItemController struct {
    itemService services.ItemService
}

func NewItemController(itemService services.ItemService) *ItemController {
    return &ItemController{itemService}
}

func (ic *ItemController) CreateItem(c echo.Context) error {
    item := new(entities.Item)
    if err := c.Bind(item); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := ic.itemService.CreateItem(item); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create item")
    }

    return c.JSON(http.StatusOK, "Item created successfully")
}

func (ic *ItemController) GetItemByID(c echo.Context) error {
    id := c.Param("id")
    item, err := ic.itemService.GetItemByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, "Item not found")
    }
    return c.JSON(http.StatusOK, item)
}

func (ic *ItemController) UpdateItem(c echo.Context) error {
    id := c.Param("id")
    item := new(entities.Item)
    if err := c.Bind(item); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := ic.itemService.UpdateItem(id, item); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to update item")
    }
    return c.JSON(http.StatusOK, "Item updated successfully")
}

func (ic *ItemController) DeleteItem(c echo.Context) error {
    id := c.Param("id")
    if err := ic.itemService.DeleteItem(id); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to delete item")
    }
    return c.JSON(http.StatusOK, "Item deleted successfully")
}