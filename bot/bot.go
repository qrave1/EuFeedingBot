package bot

import (
	"EuFeeding/repository"

	tele "gopkg.in/telebot.v4"
)

type EuFeedingBot struct {
	b *tele.Bot

	animalRepo repository.AnimalRepository
}

func NewEuFeedingBot(b *tele.Bot, animal repository.AnimalRepository) *EuFeedingBot {
	bot := &EuFeedingBot{
		b:          b,
		animalRepo: animal,
	}

	bot.InitRoutes()

	bot.initKeyboard()

	return bot
}

func (eb *EuFeedingBot) InitRoutes() {
	eb.b.Handle("/start", eb.HandleStart())
	eb.b.Handle("/help", eb.HandleHelp())

	eb.b.Handle("/add", eb.AddAnimal())
}

func (eb *EuFeedingBot) HandleStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		return c.Send("Привет!", menu)
	}
}

func (eb *EuFeedingBot) HandleHelp() func(c tele.Context) error {
	return func(c tele.Context) error {
		return c.Send("Пока пусто!")
	}
}
