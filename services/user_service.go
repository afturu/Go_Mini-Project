package services

import (
    "errors"
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
    "tukerin-platform/middleware" 
)

type UserService interface {
    Register(user *entities.User) error
    Login(email, password string) (string, uint, error)
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
    return s.userRepo.Register(user)
}

func (s *userService) Login(email, password string) (string, uint, error) {
    // Cek user berdasarkan email
    user, err := s.userRepo.Login(email, password)
    if err != nil {
        return "", 0, errors.New("invalid email or password")
    }

    // Validasi password dengan hashing
    if !middleware.CheckPasswordHash(password, user.Password) {
        return "", 0, errors.New("invalid email or password")
    }

    // Generate JWT menggunakan ID dan nama user
    token, err := s.jwtUtil.GenerateJWT(int(user.ID), user.Name)
    if err != nil {
        return "", 0, err
    }

    return token, user.ID, nil
}

func (s *userService) GetUserByID(id string) (*entities.User, error) {
    return s.userRepo.GetUserByID(id)
}

func (s *userService) UpdateUser(id string, user *entities.User) error {
    return s.userRepo.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id string) error {
    return s.userRepo.DeleteUser(id)
}