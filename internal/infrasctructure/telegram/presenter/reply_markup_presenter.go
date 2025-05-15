package presenter

import (
	"time"

	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	tele "gopkg.in/telebot.v4"
)

type ReplyMarkupPresenter struct{}

func NewReplyMarkupPresenter() *ReplyMarkupPresenter {
	return &ReplyMarkupPresenter{}
}

func (p *ReplyMarkupPresenter) MainMenu() *tele.ReplyMarkup {
	rm := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	rm.Inline(
		rm.Row(tele.Btn{Text: "Кормление", Data: "feeding:list"}),
		rm.Row(tele.Btn{Text: "Питомцы", Data: "pet:list"}),
		rm.Row(tele.Btn{Text: "Настройки", Data: "settings"}),
	)

	return rm
}

func (p *ReplyMarkupPresenter) EmptyKeyboard() *tele.ReplyMarkup {
	rm := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	rm.Inline(
		rm.Row(tele.Btn{Text: "⬅", Data: "locate:main_menu"}),
	)

	return rm
}

func (p *ReplyMarkupPresenter) CalendarKeyboard(now time.Time, highlightedDates []entity.Feeding) *tele.ReplyMarkup {
	return generateCalendarKeyboard(now, highlightedDates)
}
