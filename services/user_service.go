package services

import (
    "errors"
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
    "tukerin-platform/middleware"
)

type UserService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
    return &UserService{userRepo}
}

func (us *UserService) Register(user *entities.User) error {
    return us.userRepo.CreateUser(user)
}

func (us *UserService) Login(email, password string) (string, error) {
    user, err := us.userRepo.FindByEmail(email)
    if err != nil || user.Password != password {
        return "", errors.New("invalid credentials")
    }

    token, err := middleware.GenerateJWT(user.ID)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (us *UserService) GetUserByID(id string) (*entities.User, error) {
    return us.userRepo.GetUserByID(id)
}

func (us *UserService) UpdateUser(id string, user *entities.User) error {
    return us.userRepo.UpdateUser(id, user)
}

func (us *UserService) DeleteUser(id string) error {
    return us.userRepo.DeleteUser(id)
}