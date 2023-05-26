package usecase

import (
	"code/models"
	"code/models/payload"
	"code/repository/database"

	"github.com/labstack/echo"
)

type UserUsecase interface {
	CreateUser(req *payload.RegisterRequest) error
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(
	userRepository database.UserRepository,
) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) CreateUser(req *payload.RegisterRequest) error {
	userReq := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := u.userRepository.Createuser(userReq)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return nil
}
