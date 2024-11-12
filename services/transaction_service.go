package services

import (
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
)

type TransactionService struct {
    transactionRepo repositories.TransactionRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository) *TransactionService {
    return &TransactionService{transactionRepo}
}

func (ts *TransactionService) CreateTransaction(transaction *entities.Transaction) error {
    return ts.transactionRepo.CreateTransaction(transaction)
}

func (ts *TransactionService) GetTransactionByID(id string) (*entities.Transaction, error) {
    return ts.transactionRepo.GetTransactionByID(id)
}

func (ts *TransactionService) UpdateTransactionStatus(id, status string) error {
    return ts.transactionRepo.UpdateTransactionStatus(id, status)
}