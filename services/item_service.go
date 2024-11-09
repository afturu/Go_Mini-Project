package services

import (
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
)

type ItemService struct {
    itemRepo repositories.ItemRepository
}

func NewItemService(itemRepo repositories.ItemRepository) *ItemService {
    return &ItemService{itemRepo}
}

func (is *ItemService) CreateItem(item *entities.Item) error {
    return is.itemRepo.CreateItem(item)
}

func (is *ItemService) GetItemByID(id string) (*entities.Item, error) {
    return is.itemRepo.GetItemByID(id)
}

func (is *ItemService) UpdateItem(id string, item *entities.Item) error {
    return is.itemRepo.UpdateItem(id, item)
}

func (is *ItemService) DeleteItem(id string) error {
    return is.itemRepo.DeleteItem(id)
}