package keyboards

import (
	"emivn/internal/models"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

var (
	BtnReplenished        = tele.Btn{Text: "Пополнено"}
	BtnReplenishedAnother = tele.Btn{Text: "Пополнено на другую сумму"}
)

func Applications() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	menu.Reply(
		menu.Row(BtnActive),
		menu.Row(BtnDisputable),
		menu.Row(BtnCancel))
	return menu
}
func ApplicationList(applications []models.Application) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	var rows []tele.Row
	for _, v := range applications {
		rows = append(rows, menu.Row(tele.Btn{
			Text: fmt.Sprintf("%s / %d", v.CardID, v.Sum),
		}))
	}
	menu.Reply(rows...)
	return menu
}
func ApplicationReplenishment() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	menu.Reply(
		menu.Row(BtnReplenished),
		menu.Row(BtnReplenishedAnother),
		menu.Row(BtnCancel))
	return menu
}
