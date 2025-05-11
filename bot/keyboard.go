package bot

import tele "gopkg.in/telebot.v4"

var (
	menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	selector = &tele.ReplyMarkup{}

	listBtn = menu.Data("Список животных", "list_animal")
	//addBtn  = menu.Data("Добавить животное", "add_animal")

	btnBack = selector.Data("⬅", "back")
)

func (eb *EuFeedingBot) initKeyboard() {
	menu.Inline(
		selector.Row(listBtn, /*addBtn*/),
		selector.Row(btnBack),
	)

	eb.b.Handle(&listBtn, eb.ListAnimal())
	//eb.b.Handle(&addBtn, eb.AddAnimal())

	eb.b.Handle(&btnBack, func(c tele.Context) error {
		return c.Respond()
	})
}
