package endpoint

import (
	"emivn/internal/pkg/keyboards"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	controllerBeginState     = inputSG.New("controllerBeginState")
	controllerChoseBankState = inputSG.New("controllerChoseBankState")
	controllerInputInfoState = inputSG.New("controllerInputInfoState")
)

//func (e *Endpoint) initControllerEndpoints(manager *fsm.Manager) {
//	manager.Bind(&keyboards.BtnController, fsm.DefaultState, e.controllerBtnPressed)
//	manager.Bind(&keyboards.BtnSamuraiEnterInfo, samuraiBeginState, e.controllerOnStart)
//	manager.Bind(&keyboards.BtnTinkoff, samuraiChoseBankState, e.samuraiChoseBank)
//	manager.Bind(&keyboards.BtnSber, samuraiChoseBankState, e.samuraiChoseBank)
//	manager.Bind(tele.OnText, samuraiInputInfoState, e.samuraiInputInfo)
//}
//
//func (e *Endpoint) controllerBtnPressed(c tele.Context, state fsm.FSMContext) error {
//	if !e.serv.Repo.ControllerValid(c.Sender().Username) {
//		return c.Send("У вас нет прав")
//	}
//	state.Set(samuraiBeginState)
//	return c.Send("Выберите действие", keyboards.Controller())
//}
//
//func (e *Endpoint) controllerOnStart(c tele.Context, state fsm.FSMContext) error {
//	if time.Now().Hour() > 12 || time.Now().Hour() < 8 {
//		return c.Send("Вносить данные можно с 8:00 до 12:00")
//	}
//	state.Set(samuraiChoseBankState)
//	return c.Send("Выберите", keyboards.SamuraiChoseBankKB())
//}

func (e *Endpoint) controllerChoseBank(c tele.Context, state fsm.FSMContext) error {
	state.Update("bank", c.Message().Text)
	state.Set(samuraiInputInfoState)
	return c.Send("Введите данные")
}
func (e *Endpoint) controllerInputInfo(c tele.Context, state fsm.FSMContext) error {
	bank, _ := state.Get("bank")
	bankString := bank.(string)
	err := e.serv.SamuraiInputTurnover(c.Text(), c.Sender().Username, bankString)
	if err != nil {
		log.Println(err)
		state.Set(fsm.DefaultState)
		return c.Send("Возникла ошибка", keyboards.Default())
	}
	state.Set(samuraiBeginState)
	return c.Send("Данные сохранены")
}
