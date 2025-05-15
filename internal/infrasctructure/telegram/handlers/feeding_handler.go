package handlers

import (
	"context"
	"time"

	"github.com/qrave1/PetFeedingBot/internal/infrasctructure/telegram/presenter"
	"github.com/qrave1/PetFeedingBot/internal/usecase"
	tele "gopkg.in/telebot.v4"
)

type FeedingHandler interface {
	ShowCalendar() tele.HandlerFunc
	//CalendarCallbackHandler() tele.HandlerFunc
}

type FeedingHandlerImpl struct {
	feedingUsecase usecase.FeedingUsecase

	// TODO заменить на интерфейс
	rmPresenter *presenter.ReplyMarkupPresenter
}

func NewFeedingHandlerImpl(
	feedingUsecase usecase.FeedingUsecase,
	rmPresenter *presenter.ReplyMarkupPresenter,
) *FeedingHandlerImpl {
	return &FeedingHandlerImpl{feedingUsecase: feedingUsecase, rmPresenter: rmPresenter}
}

func (h *FeedingHandlerImpl) ShowCalendar() tele.HandlerFunc {
	return func(c tele.Context) error {
		ctx := context.Background()

		now := time.Now()

		feedings, err := h.feedingUsecase.GetForChat(ctx, c.Chat().ID, now)
		if err != nil {
			return c.Send("Не удалось получить данные о кормлении")
		}

		calendar := h.rmPresenter.CalendarKeyboard(now, feedings)

		return c.Send("График кормления", calendar)
	}
}

//func (h *FeedingHandlerImpl) CalendarCallbackHandler() tele.HandlerFunc {
//	return func(c tele.Context) error {
//
//	}
//}
