package repositories

import (
    "tukerin-platform/entities"
    "gorm.io/gorm"
)

type TransactionRepository interface {
    CreateTransaction(transaction *entities.Transaction) error
    GetTransactionByID(id string) (*entities.Transaction, error)
    UpdateTransactionStatus(id, status string) error
}

type transactionRepository struct {
    db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
    return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction *entities.Transaction) error {
    return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetTransactionByID(id string) (*entities.Transaction, error) {
    var transaction entities.Transaction
    if err := r.db.First(&transaction, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &transaction, nil
}

func (r *transactionRepository) UpdateTransactionStatus(id, status string) error {
    return r.db.Model(&entities.Transaction{}).Where("id = ?", id).Update("status", status).Error
}