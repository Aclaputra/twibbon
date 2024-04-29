package usecase

import (
	"twibbon-server/models"
	"twibbon-server/repository"
)

type TwibbonUsecase struct {
	twibbonRepository *repository.TwibbonRepository
}

type twibbonRepository interface {
	Create(twibbon models.Twibbon) (err error)
	Delete(id string) (err error)
}

func NewTwibbonUsecase(twibbonRepository *repository.TwibbonRepository) *TwibbonUsecase {
	return &TwibbonUsecase{
		twibbonRepository: twibbonRepository,
	}
}

func Create(twibbon models.Twibbon) (err error) {

	return
}

func Delete(id string) (err error) {

	return
}
