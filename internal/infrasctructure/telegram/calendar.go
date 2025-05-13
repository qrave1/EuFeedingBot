package telegram

import (
	"fmt"
	"time"

	tele "gopkg.in/telebot.v4"
)

func generateCalendarKeyboard(now time.Time, highlightedDates []time.Time) *tele.ReplyMarkup {
	rm := &tele.ReplyMarkup{}
	rows := make([]tele.Row, 0, 3)

	// TODO: добавить везде для пустых Data = "empty" и добавить эту обработку в callback handler

	year := now.Year()
	month := now.Month()

	// Навигационная строка
	prevBtn := tele.Btn{
		Unique: fmt.Sprintf("prev:%d-%02d", year, month),
		Text:   "⬅",
		Data:   fmt.Sprintf("prev:%d-%02d", year, month),
	}
	titleBtn := tele.Btn{
		Text: fmt.Sprintf("%s %d", month.String(), year),
		Data: "empty",
	}
	nextBtn := tele.Btn{
		Unique: fmt.Sprintf("next:%d-%02d", year, month),
		Text:   "➡",
		Data:   fmt.Sprintf("next:%d-%02d", year, month),
	}

	rows = append(rows, rm.Row(prevBtn, titleBtn, nextBtn))

	// Строка с названиями дней недели
	days := []string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}
	dayBtns := make([]tele.Btn, 0, 7)
	for _, day := range days {
		dayBtns = append(
			dayBtns,
			tele.Btn{
				Text: day,
				Data: fmt.Sprintf("day:%s", day),
			},
		)
	}
	rows = append(rows, rm.Row(dayBtns...))

	// Первый день месяца
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	startWeekday := int(firstDay.Weekday())
	// Количество дней в месяце
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// Создание кнопок для дат
	var buttons []tele.Btn
	// Добавление начальных заполнителей
	for i := 0; i < startWeekday; i++ {
		buttons = append(buttons, tele.Btn{Data: "empty", Text: " "})
	}
	// Добавление кнопок для дат
	for day := 1; day <= lastDay; day++ {
		text := fmt.Sprintf("%d", day)
		// Проверка, является ли дата подсвечиваемой
		for _, hl := range highlightedDates {
			if hl.Year() == year && hl.Month() == month && hl.Day() == day {
				text += " 😋"
				break
			}
		}
		callback := fmt.Sprintf("date:%d-%02d-%02d", year, month, day)
		btn := tele.Btn{
			Unique: callback,
			Text:   text,
			Data:   callback,
		}
		buttons = append(buttons, btn)
	}
	// Добавление конечных заполнителей для выравнивания
	totalButtons := len(buttons)
	placeholdersNeeded := (7 - totalButtons%7) % 7
	for i := 0; i < placeholdersNeeded; i++ {
		buttons = append(buttons, tele.Btn{Data: "empty", Text: " "})
	}
	// Группировка кнопок в строки
	for i := 0; i < len(buttons); i += 7 {
		end := i + 7

		if end > len(buttons) {
			end = len(buttons)
		}

		row := buttons[i:end]

		rows = append(rows, rm.Row(row...))
	}

	rm.Inline(rows...)

	return rm
}
