package services

// import (
// 	"testing"
// 	"tukerin-platform/entities"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mock untuk UserProfileRepository
// type MockUserProfileRepository struct {
// 	mock.Mock
// }

// func (m *MockUserProfileRepository) CreateProfile(profile *entities.UserProfile) error {
// 	args := m.Called(profile)
// 	return args.Error(0)
// }

// func (m *MockUserProfileRepository) GetProfileByUserID(userID string) (*entities.UserProfile, error) {
// 	args := m.Called(userID)
// 	return args.Get(0).(*entities.UserProfile), args.Error(1)
// }

// func (m *MockUserProfileRepository) UpdateProfile(userID string, profile *entities.UserProfile) error {
// 	args := m.Called(userID, profile)
// 	return args.Error(0)
// }

// func (m *MockUserProfileRepository) GetAllProfiles() ([]entities.UserProfile, error) {
// 	args := m.Called()
// 	return args.Get(0).([]entities.UserProfile), args.Error(1)
// }

// func (m *MockUserProfileRepository) DeleteProfile(userID string) error {
// 	args := m.Called(userID)
// 	return args.Error(0)
// }

// // Unit Test for CreateProfile
// func TestCreateProfile(t *testing.T) {
// 	mockRepo := new(MockUserProfileRepository)
// 	service := NewUserProfileService(mockRepo)

// 	profile := &entities.UserProfile{UserID: 1, Address: "Jakarta"}

// 	// Simulasi sukses
// 	mockRepo.On("CreateProfile", profile).Return(nil)

// 	err := service.CreateProfile(profile)
// 	assert.NoError(t, err)

// 	// Simulasi error
// 	mockRepo.On("CreateProfile", profile).Return(assert.AnError)

// 	err = service.CreateProfile(profile)
// 	assert.Error(t, err)
// }

// // Unit Test for GetProfileByUserID
// func TestGetProfileByUserID(t *testing.T) {
// 	mockRepo := new(MockUserProfileRepository)
// 	service := NewUserProfileService(mockRepo)

// 	profile := &entities.UserProfile{UserID: 1, Address: "Jakarta"}

// 	// Simulasi berhasil mendapatkan profile
// 	mockRepo.On("GetProfileByUserID", "123").Return(profile, nil)

// 	result, err := service.GetProfileByUserID("123")
// 	assert.NoError(t, err)
// 	assert.Equal(t, profile, result)

// 	// Simulasi gagal mendapatkan profile
// 	mockRepo.On("GetProfileByUserID", "999").Return(nil, assert.AnError)

// 	result, err = service.GetProfileByUserID("999")
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }

// // Unit Test for UpdateProfile
// func TestUpdateProfile(t *testing.T) {
// 	mockRepo := new(MockUserProfileRepository)
// 	service := NewUserProfileService(mockRepo)

// 	profile := &entities.UserProfile{UserID: 1, Address: "Jakarta"}

// 	// Simulasi sukses update profile
// 	mockRepo.On("UpdateProfile", "123", profile).Return(nil)

// 	err := service.UpdateProfile("123", profile)
// 	assert.NoError(t, err)

// 	// Simulasi gagal update profile
// 	mockRepo.On("UpdateProfile", "123", profile).Return(assert.AnError)

// 	err = service.UpdateProfile("123", profile)
// 	assert.Error(t, err)
// }

// // Unit Test for GetAllProfiles
// func TestGetAllProfiles(t *testing.T) {
// 	mockRepo := new(MockUserProfileRepository)
// 	service := NewUserProfileService(mockRepo)

// 	profile := entities.UserProfile{UserID: 1, Address: "Jakarta"}

// 	// Simulasi berhasil mendapatkan semua profile
// 	mockRepo.On("GetAllProfiles").Return([]entities.UserProfile{profile}, nil)

// 	result, err := service.GetAllProfiles()
// 	assert.NoError(t, err)
// 	assert.Len(t, result, 1)
// 	assert.Equal(t, profile, result[0])

// 	// Simulasi gagal mendapatkan profile
// 	mockRepo.On("GetAllProfiles").Return(nil, assert.AnError)

// 	result, err = service.GetAllProfiles()
// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// }

// // Unit Test for DeleteProfile
// func TestDeleteProfile(t *testing.T) {
// 	mockRepo := new(MockUserProfileRepository)
// 	service := NewUserProfileService(mockRepo)

// 	// Simulasi sukses delete profile
// 	mockRepo.On("DeleteProfile", "123").Return(nil)

// 	err := service.DeleteProfile("123")
// 	assert.NoError(t, err)

// 	// Simulasi gagal delete profile
// 	mockRepo.On("DeleteProfile", "123").Return(assert.AnError)

// 	err = service.DeleteProfile("123")
// 	assert.Error(t, err)
// }