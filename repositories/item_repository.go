package repositories

import (
    "gorm.io/gorm"
    "tukerin-platform/entities"
)

type ItemRepository interface {
    Create(item *entities.Item) error
    FindByID(id string) (*entities.Item, error)
    Update(id string, item *entities.Item) error
    Delete(id string) error
    FindAll() ([]*entities.Item, error)
}

type itemRepository struct {
    db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
    return &itemRepository{db}
}

func (r *itemRepository) Create(item *entities.Item) error {
    return r.db.Create(item).Error
}

func (r *itemRepository) FindByID(id string) (*entities.Item, error) {
    var item entities.Item
    if err := r.db.Preload("User").Preload("Category").Preload("Profile").First(&item, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *itemRepository) Update(id string, item *entities.Item) error {
    return r.db.Model(&entities.Item{}).Where("id = ?", id).Updates(item).Error
}

func (r *itemRepository) Delete(id string) error {
    return r.db.Delete(&entities.Item{}, id).Error
}

func (r *itemRepository) FindAll() ([]*entities.Item, error) {
    var items []*entities.Item
    if err := r.db.Preload("User").Preload("Category").Preload("Profile").Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}