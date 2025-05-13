package bot

import (
	"errors"
	"fmt"

	"EuFeeding/internal/domain/errs"

	tele "gopkg.in/telebot.v4"
)

func (eb *EuFeedingBot) AddAnimal() tele.HandlerFunc {
	return func(c tele.Context) error {
		args := c.Args()

		if len(args) != 1 {
			return c.Send("Неправильный формат!\nПример команды -> /add <имя>")
		}

		name := args[0]

		err := eb.animalUsecase.Add(c.Chat().ID, name)
		if err != nil {
			return c.Send("Ошибка создания животного.\nПопробуйте снова!")
		}

		return c.Send(fmt.Sprintf("Животное *%s* добавлено", name), &tele.SendOptions{ParseMode: tele.ModeMarkdownV2})
	}
}

func (eb *EuFeedingBot) ListAnimal() tele.HandlerFunc {
	return func(c tele.Context) error {
		list, err := eb.animalUsecase.List(c.Chat().ID)
		if err != nil {
			if errors.Is(err, errs.AnimalsNotFound) {
				return c.Send("У вас ещё нет животных. Добавьте их с помощью команды /add <имя>!")
			}

			return c.Send("Ошибка получения списка животных")
		}

		return c.Send(fmt.Sprintf("Список ваших животных: %v", list))
	}
}
