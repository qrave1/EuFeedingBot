package telegram

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	"github.com/qrave1/PetFeedingBot/internal/domain/errs"
	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram/handlers"
	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram/presenter"
	"github.com/qrave1/PetFeedingBot/internal/usecase"
	tele "gopkg.in/telebot.v4"
)

// TODO: убрать эту обёртку, сделать что-то типо хенделеров
type PetFeedingBot struct {
	b *tele.Bot

	// TODO: отказаться от handlers

	petUsecase   usecase.PetUsecase
	petPresenter presenter.PetPresenter
	petHandler   handlers.PetHandler

	feedingUsecase usecase.FeedingUsecase
	feedingHandler handlers.FeedingHandler

	rmPresenter *presenter.ReplyMarkupPresenter
}

func NewPetFeedingBot(
	b *tele.Bot,
	petUsecase usecase.PetUsecase,
	petPresenter presenter.PetPresenter,
	petHandler handlers.PetHandler,
	feedingUsecase usecase.FeedingUsecase,
	feedingHandler handlers.FeedingHandler,
	rmPresenter *presenter.ReplyMarkupPresenter,
) *PetFeedingBot {
	bot := &PetFeedingBot{
		b:              b,
		petUsecase:     petUsecase,
		petPresenter:   petPresenter,
		petHandler:     petHandler,
		feedingUsecase: feedingUsecase,
		feedingHandler: feedingHandler,
		rmPresenter:    rmPresenter,
	}

	bot.InitRoutes()

	return bot
}

func (pf *PetFeedingBot) InitRoutes() {
	pf.b.Handle("/start", pf.HandleStart())
	pf.b.Handle("/help", pf.HandleHelp())

	pf.b.Handle("/add", pf.petHandler.AddPet())
	//pf.b.Handle("/list", pf.petHandler.PetList())
	//
	//pf.b.Handle("/calendar", pf.feedingHandler.ShowCalendar())

	pf.b.Handle(tele.OnCallback, func(c tele.Context) error {
		//ctx := context.Background()

		callback := c.Callback()

		// Обработка пустых кнопок
		if callback.Data == "empty" {
			return c.Respond()
		}

		parts := strings.Split(callback.Data, ":")

		switch parts[0] {
		case "pet":
			return pf.handlePet(c, parts[1])
		case "feeding":
			return pf.handleFeeding(c, parts[1])
		case "locate":
			return pf.handleLocate(c, parts[1])
		case "calendar":
			return pf.handleCalendar(c, parts[1])
		case "settings":
			return c.Edit("Пока пусто!", pf.rmPresenter.EmptyKeyboard())
		}

		return nil
	})
}

func (pf *PetFeedingBot) HandleStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		return c.Send("Главное меню", pf.rmPresenter.MainMenu())
	}
}

func (pf *PetFeedingBot) HandleHelp() func(c tele.Context) error {
	return func(c tele.Context) error {
		// TODO: заполнить инфой по пользованию
		return c.Send("Пока пусто!")
	}
}

func (pf *PetFeedingBot) handlePet(c tele.Context, action string) error {
	switch action {
	case "list":
		pets, err := pf.petUsecase.List(c.Chat().ID)
		if err != nil {
			if errors.Is(err, errs.PetsNotFound) {
				return c.Send("У вас ещё нет питомцев. Добавьте их с помощью команды /add <имя>!")
			}

			return c.Send("Ошибка получения списка питомцев.\nПопробуйте снова!")
		}

		msg := pf.petPresenter.ConvertPetsList(pets)

		return c.Edit(
			fmt.Sprintf("Список ваших питомцев:\n%s", msg),
			&tele.SendOptions{ParseMode: tele.ModeMarkdownV2},
			pf.rmPresenter.EmptyKeyboard(),
		)
	default:
		slog.Error("unknown pet action", slog.Any("action", action))

		return nil
	}
}

func (pf *PetFeedingBot) handleFeeding(c tele.Context, action string) error {
	switch action {
	case "list":
		feedings, err := pf.feedingUsecase.GetForChat(context.Background(), c.Chat().ID, time.Now())
		if err != nil {
			return c.Respond(&tele.CallbackResponse{Text: "Не удалось получить данные о кормлении", ShowAlert: true})
		}

		calendar := pf.rmPresenter.CalendarKeyboard(time.Now(), feedings)

		return c.Edit("График кормления", calendar)
	default:
		slog.Error("unknown feeding action", slog.Any("action", action))

		return nil
	}
}

func (pf *PetFeedingBot) handleLocate(c tele.Context, action string) error {
	switch action {
	case "main_menu":
		return c.Edit("Главное меню", pf.rmPresenter.MainMenu())
	default:
		slog.Error("unknown locate action", slog.Any("action", action))

		return nil
	}
}

func (pf *PetFeedingBot) handleCalendar(c tele.Context, action string) error {
	parts := strings.Split(action, "_")
	if len(parts) != 2 {
		slog.Error("invalid calendar action", slog.Any("action", action))

		return c.Respond(
			&tele.CallbackResponse{
				Text:      "Ошибка в календаре",
				ShowAlert: true,
			},
		)
	}

	switch parts[0] {
	case "prev", "next":
		dateStr := parts[1]
		dateParts := strings.Split(dateStr, "-")
		if len(dateParts) != 2 {
			return c.Respond(&tele.CallbackResponse{
				Text:      "Некорректный формат месяца",
				ShowAlert: true,
			})
		}

		year, err := strconv.Atoi(dateParts[0])
		if err != nil {
			return c.Respond(&tele.CallbackResponse{
				Text:      "Некорректный год",
				ShowAlert: true,
			})
		}

		monthInt, err := strconv.Atoi(dateParts[1])
		if err != nil {
			return c.Respond(&tele.CallbackResponse{
				Text:      "Некорректный месяц",
				ShowAlert: true,
			})
		}

		currentMonth := time.Month(monthInt)
		newYear, newMonth := presenter.CalculateNewDate(year, currentMonth, parts[0] == "prev")

		feedings, err := pf.feedingUsecase.GetForChat(context.Background(), c.Chat().ID, time.Now())
		if err != nil {
			return c.Respond(&tele.CallbackResponse{
				Text:      "Не удалось получить данные о кормлении",
				ShowAlert: true,
			})
		}

		newCalendar := pf.rmPresenter.CalendarKeyboard(
			time.Date(newYear, newMonth, 0, 0, 0, 0, 0, time.UTC),
			feedings,
		)

		return c.Edit(fmt.Sprintf("Календарь за %s %d", newMonth.String(), newYear), newCalendar)

	case "date":
		dateStr := parts[1]

		parsedTime, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return c.Respond(
				&tele.CallbackResponse{
					Text:      "Неправильный формат даты",
					ShowAlert: true,
				},
			)
		}

		feedings, err := pf.feedingUsecase.GetForChat(context.Background(), c.Chat().ID, parsedTime)
		if err != nil {
			return c.Respond(
				&tele.CallbackResponse{
					Text:      "Не удалось получить данные о кормлении",
					ShowAlert: true,
				},
			)
		}

		var isExist bool
		var feedingNeededToDelete entity.Feeding

		for _, feed := range feedings {
			// если такое уже есть, значит удаляем
			if parsedTime.Equal(feed.FeedingAt) {
				isExist = true
				feedingNeededToDelete = feed
			}
		}

		if isExist {
			err = pf.feedingUsecase.Delete(context.Background(), feedingNeededToDelete.ID)
			if err != nil {
				return err
			}
		} else {
			err = pf.feedingUsecase.Add(
				context.Background(),
				feedingNeededToDelete.PetID,
				parsedTime,
				feedingNeededToDelete.FoodType,
			)
			if err != nil {
				return err
			}
		}

		feedings, err = pf.feedingUsecase.GetForChat(context.Background(), c.Chat().ID, parsedTime)
		if err != nil {
			return c.Respond(
				&tele.CallbackResponse{
					Text:      "Не удалось получить данные о кормлении",
					ShowAlert: true,
				},
			)
		}

		newCalendar := pf.rmPresenter.CalendarKeyboard(
			time.Now(),
			feedings,
		)

		return c.Edit("График кормления", newCalendar)
	default:
		slog.Error("unknown calendar action", slog.Any("action", action))

		return nil
	}
}
