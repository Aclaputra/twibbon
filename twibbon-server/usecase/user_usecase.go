package usecase

import "twibbon-server/repository"

type UserUsecase struct {
	userRepository *repository.UserRepository
}

type userusecase interface {
	Register()
	Login()
}

func NewUserUsecase(userRepository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) Register() {

	return
}

func (u *UserUsecase) Login() {

	return
}
