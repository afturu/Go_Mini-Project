package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/entities"
    "tukerin-platform/services"
)

type TransactionController struct {
    transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *TransactionController {
    return &TransactionController{transactionService}
}

func (tc *TransactionController) CreateTransaction(c echo.Context) error {
    transaction := new(entities.Transaction)
    if err := c.Bind(transaction); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := tc.transactionService.CreateTransaction(transaction); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create transaction")
    }

    return c.JSON(http.StatusOK, "Transaction created successfully")
}

func (tc *TransactionController) GetTransactionByID(c echo.Context) error {
    id := c.Param("id")
    transaction, err := tc.transactionService.GetTransactionByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, "Transaction not found")
    }
    return c.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) UpdateTransactionStatus(c echo.Context) error {
    id := c.Param("id")
    var statusData struct {
        Status string `json:"status"`
    }
    if err := c.Bind(&statusData); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := tc.transactionService.UpdateTransactionStatus(id, statusData.Status); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to update transaction status")
    }
    return c.JSON(http.StatusOK, "Transaction status updated successfully")
}