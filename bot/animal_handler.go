package bot

import (
	"fmt"
	"log"
	"time"

	"EuFeeding/entity"

	"github.com/google/uuid"
	tele "gopkg.in/telebot.v4"
)

func (eb *EuFeedingBot) AddAnimal() tele.HandlerFunc {
	return func(c tele.Context) error {
		args := c.Args()

		if len(args) != 1 {
			return c.Send("Неправильный формат!\nПример команды -> /add Мурзик")
		}

		name := args[0]

		err := eb.animalRepo.Add(
			entity.Animal{
				ID:        uuid.New().String(),
				ChatID:    c.Chat().ID,
				Name:      name,
				CreatedAt: time.Now(),
			},
		)
		if err != nil {
			return c.Send("Ошибка создания животного.\nПопробуйте снова!")
		}

		return c.Send(fmt.Sprintf("Животное *%s* добавлено", name), &tele.SendOptions{ParseMode: tele.ModeMarkdownV2})
	}
}

func (eb *EuFeedingBot) ListAnimal() tele.HandlerFunc {
	return func(c tele.Context) error {
		list, err := eb.animalRepo.List(c.Chat().ID)
		if err != nil {
			log.Println("animal list err: ", err)
			return c.Send("Ошибка получения списка животных")
		}

		if len(list) == 0 {
			return c.Send("У вас ещё нет животных. Добавьте их с помощью команды /add <имя>!")
		} else {
			return c.Send(fmt.Sprintf("Список ваших животных: %v", list))
		}
	}
}
