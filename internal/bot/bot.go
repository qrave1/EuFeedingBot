package bot

import (
	"github.com/qrave1/PetFeedingBot/internal/usecase"

	tele "gopkg.in/telebot.v4"
)

// TODO: убрать эту обёртку, сделать что-то типо хенделеров
type PetFeedingBot struct {
	b *tele.Bot

	petUsecase usecase.PetUsecase
}

func NewPetFeedingBot(b *tele.Bot, petUsecase usecase.PetUsecase) *PetFeedingBot {
	bot := &PetFeedingBot{
		b:          b,
		petUsecase: petUsecase,
	}

	bot.InitRoutes()

	bot.initKeyboard()

	return bot
}

func (pf *PetFeedingBot) InitRoutes() {
	pf.b.Handle("/start", pf.HandleStart())
	pf.b.Handle("/help", pf.HandleHelp())

	pf.b.Handle("/add", pf.AddPet())
}

func (pf *PetFeedingBot) HandleStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		return c.Send("Привет!", menu)
	}
}

func (pf *PetFeedingBot) HandleHelp() func(c tele.Context) error {
	return func(c tele.Context) error {
		return c.Send("Пока пусто!")
	}
}
