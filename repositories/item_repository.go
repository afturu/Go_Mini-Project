package repositories

import (
    "tukerin-platform/entities"
    "gorm.io/gorm"
)

type ItemRepository interface {
    CreateItem(item *entities.Item) error
    GetItemByID(id string) (*entities.Item, error)
    UpdateItem(id string, item *entities.Item) error
    DeleteItem(id string) error
}

type itemRepository struct {
    db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
    return &itemRepository{db}
}

func (r *itemRepository) CreateItem(item *entities.Item) error {
    return r.db.Create(item).Error
}

func (r *itemRepository) GetItemByID(id string) (*entities.Item, error) {
    var item entities.Item
    if err := r.db.First(&item, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *itemRepository) UpdateItem(id string, item *entities.Item) error {
    return r.db.Model(&entities.Item{}).Where("id = ?", id).Updates(item).Error
}

func (r *itemRepository) DeleteItem(id string) error {
    return r.db.Delete(&entities.Item{}, "id = ?", id).Error
}