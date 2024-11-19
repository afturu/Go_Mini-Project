package repositories

import (
    "gorm.io/gorm"
    "tukerin-platform/entities"
)

type TransactionRepository interface {
    Create(transaction *entities.Transaction) error
    FindByID(id string) (*entities.Transaction, error)
    FindAll() ([]*entities.Transaction, error)
    Update(id string, transaction *entities.Transaction) error
    Delete(id string) error
}

type transactionRepository struct {
    db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
    return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction *entities.Transaction) error {
    return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindByID(id string) (*entities.Transaction, error) {
    var transaction entities.Transaction
    if err := r.db.First(&transaction, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &transaction, nil
}

func (r *transactionRepository) FindAll() ([]*entities.Transaction, error) {
    var transactions []*entities.Transaction
    if err := r.db.Find(&transactions).Error; err != nil {
        return nil, err
    }
    return transactions, nil
}

func (r *transactionRepository) Update(id string, transaction *entities.Transaction) error {
    if id == "" {
        return gorm.ErrInvalidTransaction
    }

    return r.db.Model(&entities.Transaction{}).Where("id = ?", id).Updates(transaction).Error
}

func (r *transactionRepository) Delete(id string) error {
    if id == "" {
        return gorm.ErrInvalidTransaction
    }

    return r.db.Where("id = ?", id).Delete(&entities.Transaction{}).Error
}