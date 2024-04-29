package controller

import "twibbon-server/usecase"

type TwibbonController struct {
	twibbonUsecase *usecase.TwibbonUsecase
}

type twibbonController interface {
	Upload()
	Delete()
}

func NewTwibbonController(twibbonUsecase *usecase.TwibbonUsecase) *TwibbonController {
	return &TwibbonController{
		twibbonUsecase: twibbonUsecase,
	}
}

func (c *TwibbonController) Upload() {

	return
}

func (c *TwibbonController) Delete() {

	return
}
