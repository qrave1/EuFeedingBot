package presenter

import (
	"fmt"
	"time"

	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	tele "gopkg.in/telebot.v4"
)

func generateCalendarKeyboard(now time.Time, highlightedDates []entity.Feeding) *tele.ReplyMarkup {
	calendar := &tele.ReplyMarkup{}
	rows := make([]tele.Row, 0, 3)

	year := now.Year()
	month := now.Month()

	rows = append(rows, createNavigationHeader(month, year, calendar))

	rows = append(rows, createWeekHeader(calendar))

	rows = append(rows, createDayButtons(month, year, calendar, highlightedDates)...)

	calendar.Inline(rows...)

	return calendar
}

// createNavigationHeader создаёт навигационную строку (выбор месяца и года)
func createNavigationHeader(month time.Month, year int, calendar *tele.ReplyMarkup) tele.Row {
	callbackDate := fmt.Sprintf("%d-%d", year, month)

	prevBtn := tele.Btn{
		Text: "⬅",
		Data: fmt.Sprintf("calendar:prev_%s", callbackDate),
	}

	titleBtn := tele.Btn{
		Text: fmt.Sprintf("%s %d", month.String(), year),
		Data: "empty",
	}

	nextBtn := tele.Btn{
		Text: "➡",
		Data: fmt.Sprintf("calendar:next_%s", callbackDate),
	}

	return calendar.Row(prevBtn, titleBtn, nextBtn)
}

// createWeekHeader создаёт строку с названиями дней недели
func createWeekHeader(calendar *tele.ReplyMarkup) tele.Row {
	// Строка с названиями дней недели
	days := []string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}

	dayBtns := make([]tele.Btn, 0, 7)

	for _, day := range days {
		dayBtns = append(dayBtns,
			tele.Btn{
				Text: day,
				Data: "empty",
			},
		)
	}

	return calendar.Row(dayBtns...)
}

// createDayButtons создаёт строки с кнопками для дат
func createDayButtons(month time.Month, year int, calendar *tele.ReplyMarkup, highlightedDates []entity.Feeding) []tele.Row {
	// Первый день месяца
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	startWeekday := int(firstDay.Weekday())
	// Количество дней в месяце
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// Создание кнопок для дат
	var buttons []tele.Btn
	// Добавление начальных заполнителей
	for i := 0; i < startWeekday-1; i++ {
		buttons = append(buttons,
			tele.Btn{
				Data: "empty",
				Text: " ",
			})
	}
	// Добавление кнопок для дат
	for day := 1; day <= lastDay; day++ {
		text := fmt.Sprintf("%d", day)

		// Проверка, является ли дата подсвечиваемой
		for _, hl := range highlightedDates {
			if hl.FeedingAt.Year() == year && hl.FeedingAt.Month() == month && hl.FeedingAt.Day() == day {
				text += " 😋"
				break
			}
		}

		buttons = append(buttons,
			tele.Btn{
				Text: text,
				Data: fmt.Sprintf("calendar:date_%d-%d-%d", year, month, day),
			},
		)
	}
	// Добавление конечных заполнителей для выравнивания
	totalButtons := len(buttons)
	placeholdersNeeded := (7 - totalButtons%7) % 7
	for i := 0; i < placeholdersNeeded; i++ {
		buttons = append(buttons,
			tele.Btn{
				Data: "empty",
				Text: " ",
			},
		)
	}

	rows := make([]tele.Row, 0, 4)

	// Группировка кнопок в строки
	for i := 0; i < len(buttons); i += 7 {
		end := i + 7

		if end > len(buttons) {
			end = len(buttons)
		}

		row := buttons[i:end]

		rows = append(rows, calendar.Row(row...))
	}

	return rows
}

func CalculateNewDate(year int, month time.Month, isPrev bool) (int, time.Month) {
	if isPrev {
		if month == time.January {
			return year - 1, time.December
		}
		return year, month - 1
	}

	if month == time.December {
		return year + 1, time.January
	}
	return year, month + 1
}
