package services

// import (
// 	"testing"
// 	"tukerin-platform/entities"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mock untuk TransactionRepository
// type MockTransactionRepository struct {
// 	mock.Mock
// }

// func (m *MockTransactionRepository) Create(transaction *entities.Transaction) error {
// 	args := m.Called(transaction)
// 	return args.Error(0)
// }

// func (m *MockTransactionRepository) FindByID(id string) (*entities.Transaction, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*entities.Transaction), args.Error(1)
// }

// func (m *MockTransactionRepository) FindAll() ([]*entities.Transaction, error) {
// 	args := m.Called()
// 	return args.Get(0).([]*entities.Transaction), args.Error(1)
// }

// func (m *MockTransactionRepository) Update(id string, transaction *entities.Transaction) error {
// 	args := m.Called(id, transaction)
// 	return args.Error(0)
// }

// func (m *MockTransactionRepository) Delete(id string) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// // Unit Test for CreateTransaction
// func TestCreateTransaction(t *testing.T) {
// 	mockRepo := new(MockTransactionRepository)
// 	service := NewTransactionService(mockRepo)

// 	transaction := &entities.Transaction{ID: 1, Status: "Pending", Price: 400.000}

// 	// Simulasi sukses
// 	mockRepo.On("Create", transaction).Return(nil)

// 	err := service.CreateTransaction(transaction)
// 	assert.NoError(t, err)

// 	// Simulasi error
// 	mockRepo.On("Create", transaction).Return(assert.AnError)

// 	err = service.CreateTransaction(transaction)
// 	assert.Error(t, err)
// }

// // Unit Test for GetTransactionByID
// func TestGetTransactionByID(t *testing.T) {
// 	mockRepo := new(MockTransactionRepository)
// 	service := NewTransactionService(mockRepo)

// 	transaction := &entities.Transaction{ID: 1, Status: "Pending", Price: 400.000}

// 	// Simulasi berhasil mendapatkan transaksi
// 	mockRepo.On("FindByID", "1").Return(transaction, nil)

// 	result, err := service.GetTransactionByID("1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, transaction, result)

// 	// Simulasi gagal mendapatkan transaksi
// 	mockRepo.On("FindByID", "2").Return(nil, assert.AnError)

// 	result, err = service.GetTransactionByID("2")
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }

// // Unit Test for GetAllTransactions
// func TestGetAllTransactions(t *testing.T) {
// 	mockRepo := new(MockTransactionRepository)
// 	service := NewTransactionService(mockRepo)

// 	transaction := &entities.Transaction{ID: 1, Status: "Pending", Price: 400.000}

// 	// Simulasi berhasil mendapatkan semua transaksi
// 	mockRepo.On("FindAll").Return([]*entities.Transaction{transaction}, nil)

// 	result, err := service.GetAllTransactions()
// 	assert.NoError(t, err)
// 	assert.Len(t, result, 1)
// 	assert.Equal(t, transaction, result[0])

// 	// Simulasi gagal mendapatkan transaksi
// 	mockRepo.On("FindAll").Return(nil, assert.AnError)

// 	result, err = service.GetAllTransactions()
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }

// // Unit Test for UpdateTransaction
// func TestUpdateTransaction(t *testing.T) {
// 	mockRepo := new(MockTransactionRepository)
// 	service := NewTransactionService(mockRepo)

// 	transaction := &entities.Transaction{ID: 1, Status: "Pending", Price: 400.000}

// 	// Simulasi sukses update transaksi
// 	mockRepo.On("Update", "1", transaction).Return(nil)

// 	err := service.UpdateTransaction("1", transaction)
// 	assert.NoError(t, err)

// 	// Simulasi gagal update transaksi
// 	mockRepo.On("Update", "1", transaction).Return(assert.AnError)

// 	err = service.UpdateTransaction("1", transaction)
// 	assert.Error(t, err)
// }

// // Unit Test for DeleteTransaction
// func TestDeleteTransaction(t *testing.T) {
// 	mockRepo := new(MockTransactionRepository)
// 	service := NewTransactionService(mockRepo)

// 	// Simulasi sukses delete transaksi
// 	mockRepo.On("Delete", "1").Return(nil)

// 	err := service.DeleteTransaction("1")
// 	assert.NoError(t, err)

// 	// Simulasi gagal delete transaksi
// 	mockRepo.On("Delete", "1").Return(assert.AnError)

// 	err = service.DeleteTransaction("1")
// 	assert.Error(t, err)
// }