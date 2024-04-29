package controller

import (
	"twibbon-server/usecase"
)

type UserController struct {
	userUsecase *usecase.UserUsecase
}

type userController interface {
	Register()
	Login()
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) Register() {

	return
}

func (c *UserController) Login() {

	return
}
