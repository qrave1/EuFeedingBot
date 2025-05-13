package bot

import (
	"errors"
	"fmt"

	"EuFeeding/internal/domain/errs"

	tele "gopkg.in/telebot.v4"
)

func (eb *EuFeedingBot) AddPet() tele.HandlerFunc {
	return func(c tele.Context) error {
		args := c.Args()

		if len(args) != 1 {
			return c.Send("Неправильный формат!\nПример команды -> /add <имя>")
		}

		name := args[0]

		err := eb.petUsecase.Add(c.Chat().ID, name)
		if err != nil {
			return c.Send("Ошибка создания питомца.\nПопробуйте снова!")
		}

		return c.Send(fmt.Sprintf("Питомец *%s* добавлено", name), &tele.SendOptions{ParseMode: tele.ModeMarkdownV2})
	}
}

func (eb *EuFeedingBot) PetList() tele.HandlerFunc {
	return func(c tele.Context) error {
		list, err := eb.petUsecase.List(c.Chat().ID)
		if err != nil {
			if errors.Is(err, errs.PetsNotFound) {
				return c.Send("У вас ещё нет питомцев. Добавьте их с помощью команды /add <имя>!")
			}

			return c.Send("Ошибка получения списка питомцев.\nПопробуйте снова!")
		}

		return c.Send(fmt.Sprintf("Список ваших питомцев: %v", list))
	}
}
