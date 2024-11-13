package services

import (
    "errors"
    "time"
    "github.com/dgrijalva/jwt-go"
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
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
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{userRepo}
}

func (s *userService) Register(user *entities.User) error {
    return s.userRepo.CreateUser(user)
}

func (s *userService) Login(email, password string) (string, error) {
    user, err := s.userRepo.FindByEmail(email)
    if err != nil || user.Password != password {
        return "", errors.New("invalid email or password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
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