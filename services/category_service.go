package services

import (
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
)

type CategoryService struct {
    categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) *CategoryService {
    return &CategoryService{categoryRepo}
}

func (s *CategoryService) CreateCategory(category *entities.Category) error {
    return s.categoryRepo.Create(category)
}

func (s *CategoryService) GetCategoryByID(id string) (*entities.Category, error) {
    return s.categoryRepo.FindByID(id)
}

func (s *CategoryService) UpdateCategory(id string, category *entities.Category) error {
    return s.categoryRepo.Update(id, category)
}

func (s *CategoryService) DeleteCategory(id string) error {
    return s.categoryRepo.Delete(id)
}

func (s *CategoryService) GetAllCategories() ([]*entities.Category, error) {
    return s.categoryRepo.FindAll()
}