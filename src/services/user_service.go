package services

import (
	"errors"
	"rest-app-pos/src/config"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/models"
	"rest-app-pos/src/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(payload dto.AuthRequestRegisterDto) error
	Login(payload dto.AuthRequestLoginDto) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(payload dto.AuthRequestRegisterDto) error {
	existingUser, err := s.repo.FindByEmail(payload.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
		RoleID:   2,
	}

	return s.repo.Create(user)
}

func (s *userService) Login(payload dto.AuthRequestLoginDto) (string, error) {
	user, err := s.repo.FindByEmail(payload.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.RoleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := config.AppConfig.JWT.Secret
	if secret == "" {
		secret = "defaultsecret"
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
