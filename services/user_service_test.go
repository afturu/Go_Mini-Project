package services

// import (
// 	"errors"
// 	"testing"
// 	"tukerin-platform/entities"
// 	"tukerin-platform/middleware"
// 	"tukerin-platform/repositories"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mock untuk UserRepository
// type MockUserRepository struct {
// 	mock.Mock
// }

// func (m *MockUserRepository) Register(user *entities.User) error {
// 	args := m.Called(user)
// 	return args.Error(0)
// }

// func (m *MockUserRepository) Login(email, password string) (*entities.User, error) {
// 	args := m.Called(email, password)
// 	return args.Get(0).(*entities.User), args.Error(1)
// }

// func (m *MockUserRepository) GetUserByID(id string) (*entities.User, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*entities.User), args.Error(1)
// }

// func (m *MockUserRepository) UpdateUser(id string, user *entities.User) error {
// 	args := m.Called(id, user)
// 	return args.Error(0)
// }

// func (m *MockUserRepository) DeleteUser(id string) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// // Mock untuk JwtUsers
// type MockJwtUsers struct {
// 	mock.Mock
// }

// func (m *MockJwtUsers) GenerateJWT(userID int, userName string) (string, error) {
// 	args := m.Called(userID, userName)
// 	return args.String(0), args.Error(1)
// }

// // Pastikan MockJwtUsers mengimplementasikan middleware.JwtUsers
// var _ middleware.JwtUsers = (*MockJwtUsers)(nil)

// // Unit Test untuk Register
// func TestRegister(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	mockJwt := new(MockJwtUsers)
// 	service := &userService{
// 		userRepo: mockRepo,
// 		jwtUtil:  mockJwt,  // pastikan ini mengarah ke instance mock
// 	}

// 	user := &entities.User{Email: "test@example.com", Password: "password123", Name: "John Doe"}

// 	// Simulasi sukses
// 	mockRepo.On("Register", user).Return(nil)

// 	err := service.Register(user)
// 	assert.NoError(t, err)

// 	// Simulasi error
// 	mockRepo.On("Register", user).Return(errors.New("registration failed"))

// 	err = service.Register(user)
// 	assert.Error(t, err)
// 	assert.Equal(t, "registration failed", err.Error())
// }

// // Unit Test untuk Login
// func TestLogin(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	mockJwt := new(MockJwtUsers)
// 	service := &userService{
// 		userRepo: mockRepo,
// 		jwtUtil:  mockJwt,  // pastikan ini mengarah ke instance mock
// 	}

// 	user := &entities.User{ID: 1, Email: "test@example.com", Password: "hashedPassword", Name: "John Doe"}

// 	// Simulasi validasi sukses
// 	mockRepo.On("Login", "test@example.com", "password123").Return(user, nil)
// 	mockJwt.On("GenerateJWT", 1, "John Doe").Return("valid_token", nil)

// 	// Mock CheckPasswordHash
// 	middleware.CheckPasswordHash = func(password, hash string) bool {
// 		return password == "password123"
// 	}

// 	token, userID, err := service.Login("test@example.com", "password123")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "valid_token", token)
// 	assert.Equal(t, uint(1), userID)

// 	// Simulasi invalid password
// 	mockRepo.On("Login", "test@example.com", "wrongpassword").Return(nil, errors.New("invalid email or password"))
// 	token, userID, err = service.Login("test@example.com", "wrongpassword")
// 	assert.Error(t, err)
// 	assert.Equal(t, "", token)
// 	assert.Equal(t, uint(0), userID)
// }

// // Unit Test untuk GetUserByID
// func TestGetUserByID(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	service := &userService{
// 		userRepo: mockRepo,
// 		jwtUtil:  mockJwt,  // pastikan ini mengarah ke instance mock
// 	}

// 	user := &entities.User{ID: 1, Email: "test@example.com", Name: "John Doe"}

// 	// Simulasi berhasil mendapatkan user
// 	mockRepo.On("GetUserByID", "1").Return(user, nil)

// 	result, err := service.GetUserByID("1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, user, result)

// 	// Simulasi gagal mendapatkan user
// 	mockRepo.On("GetUserByID", "2").Return(nil, errors.New("user not found"))

// 	result, err = service.GetUserByID("2")
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }

// // Unit Test untuk UpdateUser
// func TestUpdateUser(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	service := &userService{
// 		userRepo: mockRepo,
// 		jwtUtil:  mockJwt,  // pastikan ini mengarah ke instance mock
// 	}

// 	user := &entities.User{Email: "test@example.com", Name: "Updated Name"}

// 	// Simulasi sukses update user
// 	mockRepo.On("UpdateUser", "1", user).Return(nil)

// 	err := service.UpdateUser("1", user)
// 	assert.NoError(t, err)

// 	// Simulasi gagal update user
// 	mockRepo.On("UpdateUser", "1", user).Return(errors.New("failed to update user"))

// 	err = service.UpdateUser("1", user)
// 	assert.Error(t, err)
// 	assert.Equal(t, "failed to update user", err.Error())
// }

// // Unit Test untuk DeleteUser
// func TestDeleteUser(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	service := &userService{
// 		userRepo: mockRepo,
// 		jwtUtil:  mockJwt,  // pastikan ini mengarah ke instance mock
// 	}

// 	// Simulasi sukses delete user
// 	mockRepo.On("DeleteUser", "1").Return(nil)

// 	err := service.DeleteUser("1")
// 	assert.NoError(t, err)

// 	// Simulasi gagal delete user
// 	mockRepo.On("DeleteUser", "1").Return(errors.New("failed to delete user"))

// 	err = service.DeleteUser("1")
// 	assert.Error(t, err)
// 	assert.Equal(t, "failed to delete user", err.Error())
// }