package endpoint

import (
	"emivn/internal/pkg/keyboards"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	samuraiBeginState     = inputSG.New("samuraiBeginState")
	samuraiChoseBankState = inputSG.New("samuraiChoseBankState")
	samuraiInputInfoState = inputSG.New("samuraiInputInfoState")
)

func (e *Endpoint) initSamuraiEndpoints(manager *fsm.Manager) {
	manager.Bind(&keyboards.BtnSamurai, fsm.DefaultState, e.samuraiBtnPressed)
	manager.Bind(&keyboards.BtnSamuraiEnterInfo, samuraiBeginState, e.samuraiOnStart)
	manager.Bind(&keyboards.BtnTinkoff, samuraiChoseBankState, e.samuraiChoseBank)
	manager.Bind(&keyboards.BtnSber, samuraiChoseBankState, e.samuraiChoseBank)
	manager.Bind(tele.OnText, samuraiInputInfoState, e.samuraiInputInfo)
}

func (e *Endpoint) samuraiBtnPressed(c tele.Context, state fsm.FSMContext) error {
	if !e.serv.Repo.SamuraiValid(c.Sender().Username) {
		return c.Send("У вас нет прав")
	}
	state.Set(samuraiBeginState)
	return c.Send("Выберите действие", keyboards.SamuraiKB())
}

func (e *Endpoint) samuraiOnStart(c tele.Context, state fsm.FSMContext) error {
	state.Set(samuraiChoseBankState)
	return c.Send("Выберите", keyboards.SamuraiChoseBankKB())
}

func (e *Endpoint) samuraiChoseBank(c tele.Context, state fsm.FSMContext) error {
	state.Update("bank", c.Message().Text)
	state.Set(samuraiInputInfoState)
	return c.Send("Введите данные")
}
func (e *Endpoint) samuraiInputInfo(c tele.Context, state fsm.FSMContext) error {
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
