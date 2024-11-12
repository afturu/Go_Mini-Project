package repositories

import (
    "tukerin-platform/entities"
    "gorm.io/gorm"
)

type CategoryRepository interface {
    CreateCategory(category *entities.Category) error
    GetAllCategories() ([]entities.Category, error)
}

type categoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *entities.Category) error {
    return r.db.Create(category).Error
}

func (r *categoryRepository) GetAllCategories() ([]entities.Category, error) {
    var categories []entities.Category
    if err := r.db.Find(&categories).Error; err != nil {
        return nil, err
    }
    return categories, nil
}