package bot

import (
	"EuFeeding/internal/usecase"

	tele "gopkg.in/telebot.v4"
)

type EuFeedingBot struct {
	b *tele.Bot

	animalUsecase usecase.AnimalUsecase
}

func NewEuFeedingBot(b *tele.Bot, animalUsecase usecase.AnimalUsecase) *EuFeedingBot {
	bot := &EuFeedingBot{
		b:             b,
		animalUsecase: animalUsecase,
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
