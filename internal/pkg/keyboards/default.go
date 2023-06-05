package keyboards

import (
	tele "gopkg.in/telebot.v3"
)

var (
	defaultMenu     = &tele.ReplyMarkup{ResizeKeyboard: true, RemoveKeyboard: true}
	BtnAdmin        = tele.Btn{Text: "Админ👑"}
	BtnShogun       = tele.Btn{Text: "Сегун"}
	BtnDaimyo       = tele.Btn{Text: "Дайме"}
	BtnSamurai      = tele.Btn{Text: "Самурай🥷"}
	BtnMainOperator = tele.Btn{Text: "Главный оператор"}
	BtnCollector    = tele.Btn{Text: "Инкассатор💵"}
	BtnController   = tele.Btn{Text: "Контроллер"}
	BtnCancel       = tele.Btn{Text: "❌ Отмена"}
	BtnActive       = tele.Btn{Text: "Активные"}
	BtnDisputable   = tele.Btn{Text: "Спорные"}
)

func Default() *tele.ReplyMarkup {
	defaultMenu.Reply(
		defaultMenu.Row(BtnAdmin),
		defaultMenu.Row(BtnShogun),
		defaultMenu.Row(BtnDaimyo),
		defaultMenu.Row(BtnSamurai),
		defaultMenu.Row(BtnMainOperator),
		defaultMenu.Row(BtnCollector),
		defaultMenu.Row(BtnController),
		defaultMenu.Row(BtnCancel),
	)
	return defaultMenu
}
