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

func (s *ItemService) CreateItem(item *entities.Item) error {
    return s.itemRepo.Create(item)
}

func (s *ItemService) GetItemByID(id string) (*entities.Item, error) {
    return s.itemRepo.FindByID(id)
}

func (s *ItemService) UpdateItem(id string, item *entities.Item) error {
    return s.itemRepo.Update(id, item)
}

func (s *ItemService) DeleteItem(id string) error {
    return s.itemRepo.Delete(id)
}

func (s *ItemService) GetAllItems() ([]*entities.Item, error) {
    return s.itemRepo.FindAll()
}