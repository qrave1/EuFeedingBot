package telegram

import tele "gopkg.in/telebot.v4"

var (
	menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	selector = &tele.ReplyMarkup{}

	listBtn = menu.Data("Список питомцев", "list_pet")
	//addBtn  = menu.Data("Добавить питомца", "add_pet")

	btnBack = selector.Data("⬅", "back")
)

func (pf *PetFeedingBot) initKeyboard() {
	menu.Inline(
		selector.Row(listBtn /*addBtn*/),
		//selector.Row(btnBack),
	)

	//pf.b.Handle(&listBtn, pf.PetList())
	//eb.b.Handle(&addBtn, eb.AddPet())

	//eb.b.Handle(&btnBack, func(c tele.Context) error {
	//	return c.Respond()
	//})
}
