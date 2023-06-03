package keyboards

import (
	"emivn/internal/models"
	tele "gopkg.in/telebot.v3"
)

func CardList(cards []models.Card) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	var rows []tele.Row
	for _, v := range cards {
		rows = append(rows, menu.Row(tele.Btn{
			Text: v.ID,
		}))
	}
	menu.Reply(rows...)
	return menu
}
