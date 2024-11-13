package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tukerin-platform/services"
    "tukerin-platform/entities"
)

type TransactionController struct {
    transactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
    return &TransactionController{transactionService}
}

func (c *TransactionController) CreateTransaction(ctx echo.Context) error {
    var transaction entities.Transaction
    if err := ctx.Bind(&transaction); err != nil {
        return ctx.JSON(http.StatusBadRequest, err.Error())
    }
    if err := c.transactionService.CreateTransaction(&transaction); err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, transaction)
}

func (c *TransactionController) GetTransactionByID(ctx echo.Context) error {
    id := ctx.Param("id")
    transaction, err := c.transactionService.GetTransactionByID(id)
    if err != nil {
        return ctx.JSON(http.StatusNotFound, err.Error())
    }
    return ctx.JSON(http.StatusOK, transaction)
}

func (c *TransactionController) GetAllTransactions(ctx echo.Context) error {
    transactions, err := c.transactionService.GetAllTransactions()
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, transactions)
}