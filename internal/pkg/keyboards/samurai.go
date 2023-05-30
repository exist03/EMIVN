package keyboards

import tele "gopkg.in/telebot.v3"

var (
	samuraiMenu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnSamuraiEnterInfo = tele.Btn{Text: "Ввести данные на конец смены"}
	BtnTinkoff          = tele.Btn{Text: "Тинькофф"}
	BtnSber             = tele.Btn{Text: "Сбер"}
)

func SamuraiKB() *tele.ReplyMarkup {
	samuraiMenu.Reply(
		samuraiMenu.Row(BtnSamuraiEnterInfo),
		samuraiMenu.Row(BtnCancel))
	return samuraiMenu
}

func SamuraiChoseBankKB() *tele.ReplyMarkup {
	samuraiMenu.Reply(
		samuraiMenu.Row(BtnTinkoff),
		samuraiMenu.Row(BtnSber),
		samuraiMenu.Row(BtnCancel))
	return samuraiMenu
}
