package database

import (
	"code/config"
	"code/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(username, password string) (user *models.User, err error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (a *authRepository) Login(username, password string) (user *models.User, err error) {
	err = config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
