package services

// import (
// 	"testing"
// 	"tukerin-platform/entities"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mock untuk CategoryRepository
// type MockCategoryRepository struct {
// 	mock.Mock
// }

// func (m *MockCategoryRepository) Create(category *entities.Category) error {
// 	args := m.Called(category)
// 	return args.Error(0)
// }

// func (m *MockCategoryRepository) FindByID(id string) (*entities.Category, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*entities.Category), args.Error(1)
// }

// func (m *MockCategoryRepository) Update(id string, category *entities.Category) error {
// 	args := m.Called(id, category)
// 	return args.Error(0)
// }

// func (m *MockCategoryRepository) Delete(id string) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// func (m *MockCategoryRepository) FindAll() ([]*entities.Category, error) {
// 	args := m.Called()
// 	return args.Get(0).([]*entities.Category), args.Error(1)
// }

// // Unit Test for CreateCategory
// func TestCreateCategory(t *testing.T) {
// 	mockRepo := new(MockCategoryRepository)
// 	service := NewCategoryService(mockRepo)

// 	category := &entities.Category{Name: "Electronics"}

// 	// Simulasi sukses
// 	mockRepo.On("Create", category).Return(nil)

// 	err := service.CreateCategory(category)
// 	assert.NoError(t, err)

// 	// Simulasi error
// 	mockRepo.On("Create", category).Return(assert.AnError)

// 	err = service.CreateCategory(category)
// 	assert.Error(t, err)
// }

// // Unit Test for GetCategoryByID
// func TestGetCategoryByID(t *testing.T) {
// 	mockRepo := new(MockCategoryRepository)
// 	service := NewCategoryService(mockRepo)

// 	category := &entities.Category{ID: 1, Name: "Electronics"}

// 	// Simulasi berhasil mendapatkan kategori
// 	mockRepo.On("FindByID", "1").Return(category, nil)

// 	result, err := service.GetCategoryByID("1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, category, result)

// 	// Simulasi gagal mendapatkan kategori
// 	mockRepo.On("FindByID", "2").Return(nil, assert.AnError)

// 	result, err = service.GetCategoryByID("2")
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }

// // Unit Test for UpdateCategory
// func TestUpdateCategory(t *testing.T) {
// 	mockRepo := new(MockCategoryRepository)
// 	service := NewCategoryService(mockRepo)

// 	category := &entities.Category{ID: 1, Name: "Electronics"}

// 	// Simulasi sukses update kategori
// 	mockRepo.On("Update", "1", category).Return(nil)

// 	err := service.UpdateCategory("1", category)
// 	assert.NoError(t, err)

// 	// Simulasi gagal update kategori
// 	mockRepo.On("Update", "1", category).Return(assert.AnError)

// 	err = service.UpdateCategory("1", category)
// 	assert.Error(t, err)
// }

// // Unit Test for DeleteCategory
// func TestDeleteCategory(t *testing.T) {
// 	mockRepo := new(MockCategoryRepository)
// 	service := NewCategoryService(mockRepo)

// 	// Simulasi sukses delete kategori
// 	mockRepo.On("Delete", "1").Return(nil)

// 	err := service.DeleteCategory("1")
// 	assert.NoError(t, err)

// 	// Simulasi gagal delete kategori
// 	mockRepo.On("Delete", "1").Return(assert.AnError)

// 	err = service.DeleteCategory("1")
// 	assert.Error(t, err)
// }

// // Unit Test for GetAllCategories
// func TestGetAllCategories(t *testing.T) {
// 	mockRepo := new(MockCategoryRepository)
// 	service := NewCategoryService(mockRepo)

// 	category := &entities.Category{ID: 1, Name: "Electronics"}

// 	// Simulasi berhasil mendapatkan semua kategori
// 	mockRepo.On("FindAll").Return([]*entities.Category{category}, nil)

// 	result, err := service.GetAllCategories()
// 	assert.NoError(t, err)
// 	assert.Len(t, result, 1)
// 	assert.Equal(t, category, result[0])

// 	// Simulasi gagal mendapatkan kategori
// 	mockRepo.On("FindAll").Return(nil, assert.AnError)

// 	result, err = service.GetAllCategories()
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }