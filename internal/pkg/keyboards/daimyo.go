package keyboards

import tele "gopkg.in/telebot.v3"

var (
	daimyoMenu          = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnHierarchy        = tele.Btn{Text: "Иерархия"}
	BtnReport           = tele.Btn{Text: "Отчет"}
	BtnCardLimit        = tele.Btn{Text: "Лимит по карте"}
	BtnApplications     = tele.Btn{Text: "Заявки"}
	BtnAskReplenishment = tele.Btn{Text: "Запросить пополнение"}
	BtnCreateSamurai    = tele.Btn{Text: "Создать самурая"}
	BtnSubordinates     = tele.Btn{Text: "Список подчиненных"}
	BtnEnterInfoShift   = tele.Btn{Text: "Ввести данные за смену"}
	BtnShift            = tele.Btn{Text: "За смену"}
	BtnPeriod           = tele.Btn{Text: "За период"}
)

func Daimyo() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnAskReplenishment),
		daimyoMenu.Row(BtnApplications),
		daimyoMenu.Row(BtnCardLimit),
		daimyoMenu.Row(BtnReport),
		daimyoMenu.Row(BtnHierarchy),
		daimyoMenu.Row(BtnCancel),
	)
	return daimyoMenu
}

func DaimyoHierarchy() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnCreateSamurai),
		daimyoMenu.Row(BtnSubordinates),
		daimyoMenu.Row(BtnCancel),
	)
	return daimyoMenu
}
func DaimyoReport() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnEnterInfoShift),
		daimyoMenu.Row(BtnReport),
		daimyoMenu.Row(BtnCancel),
	)
	return daimyoMenu
}
func DaimyoReportPeriod() *tele.ReplyMarkup {
	daimyoMenu.Reply(
		daimyoMenu.Row(BtnShift),
		daimyoMenu.Row(BtnPeriod),
		daimyoMenu.Row(BtnCancel),
	)
	return daimyoMenu
}
