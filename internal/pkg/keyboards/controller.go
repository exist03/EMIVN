package keyboards

import tele "gopkg.in/telebot.v3"

var (
	controllerMenu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnControllerEnterInfo = tele.Btn{Text: "Ввести данные"}
)

func Controller() *tele.ReplyMarkup {
	controllerMenu.Reply(
		controllerMenu.Row(BtnControllerEnterInfo),
		controllerMenu.Row(BtnCancel))
	return controllerMenu
}
