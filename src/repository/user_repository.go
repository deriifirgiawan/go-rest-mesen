package repository

import (
	"rest-app-pos/src/database"
	"rest-app-pos/src/models"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}


func (r *userRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
