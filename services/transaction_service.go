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

func (s *TransactionService) CreateTransaction(transaction *entities.Transaction) error {
    return s.transactionRepo.Create(transaction)
}

func (s *TransactionService) GetTransactionByID(id string) (*entities.Transaction, error) {
    return s.transactionRepo.FindByID(id)
}

func (s *TransactionService) GetAllTransactions() ([]*entities.Transaction, error) {
    return s.transactionRepo.FindAll()
}