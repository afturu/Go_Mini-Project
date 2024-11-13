package services

import (
    "errors"
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
    "tukerin-platform/middleware" 
)

type UserService interface {
    Register(user *entities.User) error
    Login(email, password string) (string, error)
    GetUserByID(id string) (*entities.User, error)
    UpdateUser(id string, user *entities.User) error
    DeleteUser(id string) error
}

type userService struct {
    userRepo repositories.UserRepository
    jwtUtil  middleware.JwtUsers
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{
        userRepo: userRepo,
        jwtUtil:  middleware.JwtUsers{},
    }
}

func (s *userService) Register(user *entities.User) error {
    return s.userRepo.CreateUser(user)
}

func (s *userService) Login(email, password string) (string, error) {
    user, err := s.userRepo.FindByEmail(email)
    if err != nil || user.Password != password {
        return "", errors.New("invalid email or password")
    }

    // Menggunakan fungsi GenerateJWT dari middleware
    token, err := s.jwtUtil.GenerateJWT(int(user.ID), user.Name)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (s *userService) GetUserByID(id string) (*entities.User, error) {
    return s.userRepo.FindUserByID(id)
}

func (s *userService) UpdateUser(id string, user *entities.User) error {
    return s.userRepo.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id string) error {
    return s.userRepo.DeleteUser(id)
}