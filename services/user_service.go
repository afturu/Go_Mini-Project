package services

import (
    "errors"
    "tukerin-platform/entities"
    "tukerin-platform/repositories"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

type UserService struct {
    userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
    return &UserService{userRepo}
}

func (us *UserService) Register(user *entities.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return us.userRepo.Register(user)
}

func (us *UserService) Login(username, password string) (string, error) {
    user, err := us.userRepo.FindByUsername(username)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
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