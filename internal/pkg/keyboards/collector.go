package keyboards

import tele "gopkg.in/telebot.v3"

var (
	collectorMenu = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnEnterInfo  = tele.Btn{Text: "Ввести данные"}
)

func Collector() *tele.ReplyMarkup {
	collectorMenu.Reply(
		collectorMenu.Row(BtnEnterInfo),
		collectorMenu.Row(BtnCancel))
	return collectorMenu
}
