package controllers

import (
	"code/models/payload"
	"code/usecase"

	"github.com/labstack/echo"
)

type AuthController interface{}

type authController struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
}

func NewAuthController(
	userUsecase usecase.UserUsecase,
	authUsecase usecase.AuthUsecase,
) *authController {
	return &authController{
		userUsecase,
		authUsecase,
	}
}

func (a *authController) LoginController(c echo.Context) error {
	req := payload.LoginRequest{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := a.authUsecase.LoginUser(&req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success login",
		Data:    res,
	})
}

func (a *authController) RegisterController(c echo.Context) error {
	req := payload.RegisterRequest{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	err := a.userUsecase.CreateUser(&req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success create user",
		Data:    req,
	})

}
