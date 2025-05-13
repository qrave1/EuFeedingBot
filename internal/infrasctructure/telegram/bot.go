package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram/handlers"
	tele "gopkg.in/telebot.v4"
)

// TODO: убрать эту обёртку, сделать что-то типо хенделеров
type PetFeedingBot struct {
	b *tele.Bot

	petHandler handlers.PetHandler
}

func NewPetFeedingBot(b *tele.Bot, petHandler handlers.PetHandler) *PetFeedingBot {
	bot := &PetFeedingBot{
		b:          b,
		petHandler: petHandler,
	}

	bot.InitRoutes()

	bot.initKeyboard()

	return bot
}

func (pf *PetFeedingBot) InitRoutes() {
	pf.b.Handle("/start", pf.HandleStart())
	pf.b.Handle("/help", pf.HandleHelp())

	pf.b.Handle("/add", pf.petHandler.AddPet())
	pf.b.Handle("/list", pf.petHandler.PetList())

	pf.b.Handle("/calendar", func(c tele.Context) error {
		now := time.Now()

		highlighted := []time.Time{
			// Пример подсвечиваемых дат
			time.Date(now.Year(), now.Month(), 15, 0, 0, 0, 0, time.UTC),
			time.Date(now.Year(), now.Month(), 20, 0, 0, 0, 0, time.UTC),
		}

		calendar := generateCalendarKeyboard(now, highlighted)

		return c.Send("Выберите дату", calendar)
	})

	pf.b.Handle(tele.OnCallback, func(c tele.Context) error {
		callback := c.Callback()
		parts := strings.Split(callback.Data, ":")
		if len(parts) == 2 {
			action := parts[0]
			dateStr := parts[1]
			year, _ := strconv.Atoi(strings.Split(dateStr, "-")[0])
			monthInt, _ := strconv.Atoi(strings.Split(dateStr, "-")[1])
			month := time.Month(monthInt)
			var newYear int
			var newMonth time.Month
			if action == "prev" {
				if month == time.January {
					newYear = year - 1
					newMonth = time.December
				} else {
					newYear = year
					newMonth = month - 1
				}
			} else if action == "next" {
				if month == time.December {
					newYear = year + 1
					newMonth = time.January
				} else {
					newYear = year
					newMonth = month + 1
				}
			}
			if action == "prev" || action == "next" {
				highlighted := []time.Time{
					time.Date(newYear, newMonth, 15, 0, 0, 0, 0, time.UTC),
					time.Date(newYear, newMonth, 20, 0, 0, 0, 0, time.UTC),
				}
				newCalendar := generateCalendarKeyboard(time.Date(newYear, newMonth, 0, 0, 0, 0, 0, time.UTC), highlighted)
				return c.Edit(fmt.Sprintf("Календарь за %s %d", newMonth.String(), newYear), newCalendar)
			}
		} else if strings.HasPrefix(callback.Data, "date:") {
			// Обработка выбора даты
			dateStr := strings.TrimPrefix(callback.Data, "date:")
			date, err := time.Parse("2006-01-02", dateStr)
			if err == nil {
				return c.Send(fmt.Sprintf("Вы выбрали %s", date.Format("2006-01-02")))
			}
		}
		return nil
	})
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
