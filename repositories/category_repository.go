package repositories

import (
    "gorm.io/gorm"
    "tukerin-platform/entities"
)

type CategoryRepository interface {
    Create(category *entities.Category) error
    FindByID(id string) (*entities.Category, error)
    Update(id string, category *entities.Category) error
    Delete(id string) error
    FindAll() ([]*entities.Category, error)
}

type categoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *entities.Category) error {
    return r.db.Create(category).Error
}

func (r *categoryRepository) FindByID(id string) (*entities.Category, error) {
    var category entities.Category
    if err := r.db.First(&category, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &category, nil
}

func (r *categoryRepository) Update(id string, category *entities.Category) error {
    return r.db.Model(&entities.Category{}).Where("id = ?", id).Updates(category).Error
}

func (r *categoryRepository) Delete(id string) error {
    return r.db.Delete(&entities.Category{}, id).Error
}

func (r *categoryRepository) FindAll() ([]*entities.Category, error) {
    var categories []*entities.Category
    if err := r.db.Find(&categories).Error; err != nil {
        return nil, err
    }
    return categories, nil
}