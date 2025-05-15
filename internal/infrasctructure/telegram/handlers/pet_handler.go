package handlers

import (
	"fmt"

	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram/presenter"
	"github.com/qrave1/PetFeedingBot/internal/usecase"

	tele "gopkg.in/telebot.v4"
)

type PetHandler interface {
	AddPet() tele.HandlerFunc
	//PetList() tele.HandlerFunc
}

type PetHandlerImpl struct {
	petUsecase   usecase.PetUsecase
	petPresenter presenter.PetPresenter
}

func NewPetHandlerImpl(petUsecase usecase.PetUsecase, petPresenter presenter.PetPresenter) *PetHandlerImpl {
	return &PetHandlerImpl{petUsecase: petUsecase, petPresenter: petPresenter}
}

func (ph *PetHandlerImpl) AddPet() tele.HandlerFunc {
	return func(c tele.Context) error {
		args := c.Args()

		if len(args) != 1 {
			return c.Send("Неправильный формат!\nПример команды -> /add <имя>")
		}

		name := args[0]

		err := ph.petUsecase.Add(c.Chat().ID, name)
		if err != nil {
			return c.Send("Ошибка создания питомца.\nПопробуйте снова!")
		}

		return c.Send(fmt.Sprintf("Питомец *%s* добавлено", name), &tele.SendOptions{ParseMode: tele.ModeMarkdownV2})
	}
}

// TODO: cleanup
//func (ph *PetHandlerImpl) PetList() tele.HandlerFunc {
//	return func(c tele.Context) error {
//
//	}
//}
