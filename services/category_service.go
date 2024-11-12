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

func (cs *CategoryService) CreateCategory(category *entities.Category) error {
    return cs.categoryRepo.CreateCategory(category)
}

func (cs *CategoryService) GetAllCategories() ([]entities.Category, error) {
    return cs.categoryRepo.GetAllCategories()
}