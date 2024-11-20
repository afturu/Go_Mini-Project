package services

// import (
// 	"testing"
// 	"tukerin-platform/entities"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mock untuk ItemRepository
// type MockItemRepository struct {
// 	mock.Mock
// }

// func (m *MockItemRepository) Create(item *entities.Item) error {
// 	args := m.Called(item)
// 	return args.Error(0)
// }

// func (m *MockItemRepository) FindByID(id string) (*entities.Item, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*entities.Item), args.Error(1)
// }

// func (m *MockItemRepository) Update(id string, item *entities.Item) error {
// 	args := m.Called(id, item)
// 	return args.Error(0)
// }

// func (m *MockItemRepository) Delete(id string) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// func (m *MockItemRepository) FindAll() ([]*entities.Item, error) {
// 	args := m.Called()
// 	return args.Get(0).([]*entities.Item), args.Error(1)
// }

// // Unit Test for CreateItem
// func TestCreateItem(t *testing.T) {
// 	mockRepo := new(MockItemRepository)
// 	service := NewItemService(mockRepo)

// 	item := &entities.Item{ID: 1, Title: "Laptop", Price: 324.000}

// 	// Simulasi sukses
// 	mockRepo.On("Create", item).Return(nil)

// 	err := service.CreateItem(item)
// 	assert.NoError(t, err)

// 	// Simulasi error
// 	mockRepo.On("Create", item).Return(assert.AnError)

// 	err = service.CreateItem(item)
// 	assert.Error(t, err)
// }

// // Unit Test for GetItemByID
// func TestGetItemByID(t *testing.T) {
// 	mockRepo := new(MockItemRepository)
// 	service := NewItemService(mockRepo)

// 	item := &entities.Item{ID: 1, Title: "Laptop", Price: 324.000}

// 	// Simulasi berhasil mendapatkan item
// 	mockRepo.On("FindByID", "1").Return(item, nil)

// 	result, err := service.GetItemByID("1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, item, result)

// 	// Simulasi gagal mendapatkan item
// 	mockRepo.On("FindByID", "2").Return(nil, assert.AnError)

// 	result, err = service.GetItemByID("2")
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }

// // Unit Test for UpdateItem
// func TestUpdateItem(t *testing.T) {
// 	mockRepo := new(MockItemRepository)
// 	service := NewItemService(mockRepo)

// 	item := &entities.Item{ID: 1, Title: "Laptop", Price: 324.000}

// 	// Simulasi sukses update item
// 	mockRepo.On("Update", "1", item).Return(nil)

// 	err := service.UpdateItem("1", item)
// 	assert.NoError(t, err)

// 	// Simulasi gagal update item
// 	mockRepo.On("Update", "1", item).Return(assert.AnError)

// 	err = service.UpdateItem("1", item)
// 	assert.Error(t, err)
// }

// // Unit Test for DeleteItem
// func TestDeleteItem(t *testing.T) {
// 	mockRepo := new(MockItemRepository)
// 	service := NewItemService(mockRepo)

// 	// Simulasi sukses delete item
// 	mockRepo.On("Delete", "1").Return(nil)

// 	err := service.DeleteItem("1")
// 	assert.NoError(t, err)

// 	// Simulasi gagal delete item
// 	mockRepo.On("Delete", "1").Return(assert.AnError)

// 	err = service.DeleteItem("1")
// 	assert.Error(t, err)
// }

// // Unit Test for GetAllItems
// func TestGetAllItems(t *testing.T) {
// 	mockRepo := new(MockItemRepository)
// 	service := NewItemService(mockRepo)

// 	item := &entities.Item{ID: 1, Title: "Laptop", Price: 324.000}

// 	// Simulasi berhasil mendapatkan semua item
// 	mockRepo.On("FindAll").Return([]*entities.Item{item}, nil)

// 	result, err := service.GetAllItems()
// 	assert.NoError(t, err)
// 	assert.Len(t, result, 1)
// 	assert.Equal(t, item, result[0])

// 	// Simulasi gagal mendapatkan item
// 	mockRepo.On("FindAll").Return(nil, assert.AnError)

// 	result, err = service.GetAllItems()
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }