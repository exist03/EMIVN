package endpoint

import (
	"emivn/internal/models"
	"emivn/internal/pkg/keyboards"
	"fmt"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

var (
	daimyoBeginState = inputSG.New("daimyoBeginState")
	//hierarchy
	daimyoHierarchyState       = inputSG.New("daimyoHierarchyState")
	daimyoSamuraiNicknameState = inputSG.New("daimyoSamuraiNicknameState")
	daimyoSamuraiUsernameState = inputSG.New("daimyoSamuraiUsernameState")
	//report
	daimyoReportState            = inputSG.New("daimyoReportState")
	daimyoReportReportState      = inputSG.New("daimyoReportReportState")
	daimyoReportChosePeriodState = inputSG.New("daimyoReportChosePeriodState")
	daimyoReportForPeriodState   = inputSG.New("daimyoReportForPeriodState")
	daimyoReportPeriodEndState   = inputSG.New("daimyoReportPeriodEndState")
)

func (e *Endpoint) initDaimyoEndpoints(manager *fsm.Manager) {
	manager.Bind(&keyboards.BtnDaimyo, fsm.DefaultState, e.daimyo)
	//hierarchy
	manager.Bind(&keyboards.BtnHierarchy, daimyoBeginState, e.daimyoHierarchy)
	manager.Bind(&keyboards.BtnSubordinates, daimyoHierarchyState, e.daimyoSubordinates)
	manager.Bind(&keyboards.BtnCreateSamurai, daimyoHierarchyState, e.daimyoCreateSamurai)
	manager.Bind(tele.OnText, daimyoSamuraiNicknameState, e.daimyoCreateSamuraiNickname)
	manager.Bind(tele.OnText, daimyoSamuraiUsernameState, e.daimyoCreateSamuraiUsername)
	//report
	manager.Bind(&keyboards.BtnReport, daimyoBeginState, e.daimyoReport)
	manager.Bind(&keyboards.BtnReport, daimyoReportState, e.daimyoReportReport)
	manager.Bind(&keyboards.BtnShift, daimyoReportChosePeriodState, e.daimyoReportReportShift)
	manager.Bind(&keyboards.BtnPeriod, daimyoReportChosePeriodState, e.daimyoReportReportPeriodStart)
	manager.Bind(&keyboards.BtnPeriod, daimyoReportForPeriodState, e.daimyoReportReportPeriodStartInput)
	manager.Bind(&keyboards.BtnPeriod, daimyoReportPeriodEndState, e.daimyoReportReportPeriodEndInput)

}
func (e *Endpoint) daimyo(c tele.Context, state fsm.FSMContext) error {
	if !e.serv.Repo.ValidUser(c.Sender().Username, "Daimyo") {
		return c.Send("У вас нет прав")
	}
	state.Set(daimyoBeginState)
	return c.Send("Выберите действие", keyboards.Daimyo())
}

// hierarchy
func (e *Endpoint) daimyoHierarchy(c tele.Context, state fsm.FSMContext) error {
	state.Set(daimyoHierarchyState)
	return c.Send("Выберите", keyboards.DaimyoHierarchy())
}
func (e *Endpoint) daimyoSubordinates(c tele.Context, state fsm.FSMContext) error {
	samuraiList, err := e.serv.Repo.SamuraiGetListByOwner(c.Sender().Username)
	if err != nil {
		log.Println(err)
		state.Set(fsm.DefaultState)
		return c.Send("Возникла ошбика", keyboards.Default())
	}
	for _, v := range samuraiList {
		err := c.Send(fmt.Sprintf("%s", v))
		if err != nil {
			log.Println(err)
			return c.Send("Ошибка")
		}
	}
	return c.Send("Выберите", keyboards.DaimyoHierarchy())
}
func (e *Endpoint) daimyoCreateSamurai(c tele.Context, state fsm.FSMContext) error {
	state.Set(daimyoSamuraiNicknameState)
	return c.Send("Введите имя")
}
func (e *Endpoint) daimyoCreateSamuraiNickname(c tele.Context, state fsm.FSMContext) error {
	state.Set(daimyoSamuraiUsernameState)
	state.Update("nickname", c.Text())
	return c.Send("Выберите тэг")
}
func (e *Endpoint) daimyoCreateSamuraiUsername(c tele.Context, state fsm.FSMContext) error {
	nickname := state.MustGet("nickname")
	nick := nickname.(string)
	username := c.Text()
	s := models.Samurai{
		Username: username,
		Nickname: nick,
		Owner:    c.Sender().Username,
	}
	err := e.serv.Repo.SamuraiInsert(s)
	if err != nil {
		log.Println(err)
		return c.Send("Возникла ошибка, попробуйте еще раз")
	}
	state.Set(daimyoHierarchyState)
	return c.Send("Данные записаны")
}

// report
func (e *Endpoint) daimyoReport(c tele.Context, state fsm.FSMContext) error {
	state.Set(daimyoReportState)
	return c.Send("Введите данные на конец смены с 8 до 12", keyboards.DaimyoReport())
}
func (e *Endpoint) daimyoReportReport(c tele.Context, state fsm.FSMContext) error {
	state.Set(daimyoReportChosePeriodState)
	return c.Send("Выберите", keyboards.DaimyoReportPeriod())
}
func (e *Endpoint) daimyoReportReportShift(c tele.Context, state fsm.FSMContext) error {
	mapControll, err := e.serv.Repo.ControllerGetSamuraiAmountList(c.Sender().Username)
	mapSamurai, err := e.serv.Repo.DaimyoGetSamuraiAmountList(c.Sender().Username)
	if err != nil {
		log.Println(err)
		return c.Send("Что-то пошло не так")
	}
	shiftTime := time.Now().Add(-24 * time.Hour).Format("02.01.2006")
	for k, _ := range mapSamurai {
		if mapSamurai[k].Sber != mapControll[k].Sber || mapSamurai[k].Tinkoff != mapControll[k].Tinkoff {
			amountControll := mapControll[k].Sber + mapControll[k].Tinkoff
			amountSamurai := mapSamurai[k].Sber + mapSamurai[k].Tinkoff
			c.Send(fmt.Sprintf("%s\n%s\nВсего\n +"+
				"%.0f / %.0f / %.0f\n\n "+
				"сбер\n%.0f / %.0f / %.0f\n\n"+
				"тинькофф\n%.0f / %.0f / %.0f", shiftTime, k, amountSamurai, amountControll, amountControll-amountSamurai, mapSamurai[k].Sber, mapControll[k].Sber, mapControll[k].Sber-mapSamurai[k].Sber, mapSamurai[k].Tinkoff, mapControll[k].Tinkoff, mapControll[k].Tinkoff-mapSamurai[k].Tinkoff))
		} else {
			c.Send(fmt.Sprintf("%s\nрасхождение по %s отсутствуют", shiftTime, k))
		}
	}
	return nil
}
func (e *Endpoint) daimyoReportReportPeriodStart(c tele.Context, state fsm.FSMContext) error {
	state.Set(daimyoReportForPeriodState)
	return c.Send("Введите дату начала периода")
}
func (e *Endpoint) daimyoReportReportPeriodStartInput(c tele.Context, state fsm.FSMContext) error {
	state.Set(daimyoReportPeriodEndState)
	state.Update("beginDate", c.Text())
	return c.Send("Введите дату конца периода включительно")
}
func (e *Endpoint) daimyoReportReportPeriodEndInput(c tele.Context, state fsm.FSMContext) error {
	endDate, err := time.Parse(c.Text(), "2006-01-02")
	if err != nil {
		c.Send("Некорретный ввод, попробуйте еще раз")
	}
	tmp := state.MustGet("beginDate")
	beginDate, _ := time.Parse(tmp.(string), "2006-01-02")

}
