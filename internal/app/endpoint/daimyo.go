package endpoint

import (
	"emivn/internal/models"
	"emivn/internal/pkg/keyboards"
	"fmt"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	daimyoBeginState = inputSG.New("daimyoBeginState")
	//hierarchy
	daimyoHierarchyState       = inputSG.New("daimyoHierarchyState")
	daimyoSamuraiNicknameState = inputSG.New("daimyoSamuraiNicknameState")
	daimyoSamuraiUsernameState = inputSG.New("daimyoSamuraiUsernameState")
	//report
	daimyoReportState                 = inputSG.New("daimyoReportState")
	daimyoReportChosePeriodState      = inputSG.New("daimyoReportChosePeriodState")
	daimyoReportForPeriodState        = inputSG.New("daimyoReportForPeriodState")
	daimyoReportPeriodEndState        = inputSG.New("daimyoReportPeriodEndState")
	daimyoReportEnterInfoChoseCard    = inputSG.New("daimyoReportEnterInfoChoseCard")
	daimyoReportEnterInfoEnterBalance = inputSG.New("daimyoReportEnterInfoEnterBalance")
	//create application
	daimyoCreateApplicationChoseCardState   = inputSG.New("daimyoCreateApplicationChoseCardState")
	daimyoCreateApplicationEnterAmountState = inputSG.New("daimyoCreateApplicationEnterAmountState")
	//applications
	daimyoApplicationDisputableCardState = inputSG.New("daimyoApplicationDisputableCardState")
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
	manager.Bind(tele.OnText, daimyoReportForPeriodState, e.daimyoReportReportPeriodStartInput)
	manager.Bind(tele.OnText, daimyoReportPeriodEndState, e.daimyoReportReportPeriodEndInput)
	manager.Bind(&keyboards.BtnEnterInfoShift, daimyoReportState, e.daimyoReportEnterInfo)
	manager.Bind(tele.OnText, daimyoReportEnterInfoChoseCard, e.daimyoReportEnterInfoChoseCard)
	manager.Bind(tele.OnText, daimyoReportEnterInfoEnterBalance, e.daimyoReportEnterInfoFinally)
	//card limit
	manager.Bind(&keyboards.BtnCardLimit, daimyoBeginState, e.daimyoCardLimit)
	//create application
	manager.Bind(&keyboards.BtnAskReplenishment, daimyoBeginState, e.daimyoCreateApplication)
	manager.Bind(tele.OnText, daimyoCreateApplicationChoseCardState, e.daimyoCreateApplicationChoseCard)
	manager.Bind(tele.OnText, daimyoCreateApplicationEnterAmountState, e.daimyoCreateApplicationEnterAmount)
	//Applications
	manager.Bind(&keyboards.BtnApplications, daimyoBeginState, e.daimyoApplications)
	manager.Bind(&keyboards.BtnActive, daimyoBeginState, e.daimyoApplicationsActive)
	manager.Bind(&keyboards.BtnDisputable, daimyoBeginState, e.daimyoApplicationsDisputable)
	manager.Bind(tele.OnText, daimyoApplicationDisputableCardState, e.daimyoApplicationsDisputableChoseReplenishment)
	manager.Bind(&keyboards.BtnReplenishedAnother, daimyoApplicationDisputableCardState, nil)
	manager.Bind(&keyboards.BtnReplenished, daimyoApplicationDisputableCardState, e.daimyoApplicationsDisputableChoseReplenishmentDone)

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
	nick := state.MustGet("nickname").(string)
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
	endDate, err := time.Parse("2006-01-02", c.Text())
	if err != nil {
		log.Println(err)
		state.Set(daimyoReportForPeriodState)
		c.Send("Некорретный ввод, попробуйте еще раз\nВведите дату начала периода")
		return err
	}
	tmp := state.MustGet("beginDate")
	beginDate, _ := time.Parse("2006-01-02", tmp.(string))
	mapRes, err := e.serv.Repo.DaimyoGetReportByPeriod(c.Sender().Username, beginDate, endDate)
	var result float64
	if err != nil {
		log.Println(err)
		return c.Send("Что-то пошло не так повторите попытку")
	}
	for _, v := range mapRes {
		result += v.End - v.Begin
	}
	c.Send(fmt.Sprintf("%s\nОборот: %.0f\n0.0015 -> %.0f", c.Sender().Username, result, result*0.0015))
	for k, v := range mapRes {
		c.Send(fmt.Sprintf("%s\n"+
			"%.0f / %.0f ", k, v.Begin, v.End))
	}
	state.Set(daimyoBeginState)
	return c.Send("Конец отчета", keyboards.Daimyo())
}
func (e *Endpoint) daimyoReportEnterInfo(c tele.Context, state fsm.FSMContext) error {
	cards, err := e.serv.Repo.CardGetListByOwner(c.Sender().Username)
	if err != nil {
		log.Println(err)
		return c.Send("Что-то пошло не так")
	}
	state.Set(daimyoReportEnterInfoChoseCard)
	return c.Send("Выберите", keyboards.CardList(cards))
}
func (e *Endpoint) daimyoReportEnterInfoChoseCard(c tele.Context, state fsm.FSMContext) error {
	state.Update("cardID", c.Text())
	state.Set(daimyoReportEnterInfoEnterBalance)
	return c.Send("Введите остаток баланса карты на 8:00")
}
func (e *Endpoint) daimyoReportEnterInfoFinally(c tele.Context, state fsm.FSMContext) error {
	card := state.MustGet("cardID").(string)
	amount, err := strconv.Atoi(c.Text())
	if err != nil {
		return c.Send("Некорректное значение. Попробуйте еще раз")
	}
	err = e.serv.Repo.CardUpdateBalance(card, amount)
	if err != nil {
		log.Println(err)
		state.Set(daimyoBeginState)
		return c.Send("Возникла ошибка", keyboards.Daimyo())
	}
	state.Set(daimyoBeginState)
	return c.Send("Данные записаны", keyboards.Daimyo())
}

// card limit
func (e *Endpoint) daimyoCardLimit(c tele.Context, state fsm.FSMContext) error {
	cards, err := e.serv.Repo.CardGetListByOwner(c.Sender().Username)
	if err != nil {
		log.Println(err)
		return c.Send("Что-то пошло не так")
	}
	for _, v := range cards {
		c.Send(fmt.Sprintf("%s - %d", v.ID, v.Limit))
	}
	return nil
}

// create application
func (e *Endpoint) daimyoCreateApplication(c tele.Context, state fsm.FSMContext) error {
	cards, err := e.serv.Repo.CardGetListByOwner(c.Sender().Username)
	if err != nil {
		log.Println(err)
		return c.Send("Что-то пошло не так, повториты попытку", keyboards.Daimyo())
	}
	state.Set(daimyoCreateApplicationChoseCardState)
	return c.Send("Выберите карту", keyboards.CardList(cards))
}
func (e *Endpoint) daimyoCreateApplicationChoseCard(c tele.Context, state fsm.FSMContext) error {
	cardID := c.Text()
	if e.serv.CardInDispute(cardID) {
		return c.Send("Вы не можете использовать данную карту так как она активна/в споре")
	}
	state.Update("cardID", cardID)
	state.Set(daimyoCreateApplicationEnterAmountState)
	return c.Send("Введите сумму")
}
func (e *Endpoint) daimyoCreateApplicationEnterAmount(c tele.Context, state fsm.FSMContext) error {
	cardID := state.MustGet("cardID").(string)
	amount, err := strconv.Atoi(c.Text())
	if err != nil {
		return c.Send("Введите корректное значение")
	}
	card, err := e.serv.CardGetByID(cardID)
	if err != nil {
		log.Println(err)
		return c.Send("Произошла ошибка\nПопробуйте еще раз")
	}
	if amount > card.Limit {
		return c.Send(fmt.Sprintf("Вы превысили лимит\nОстаток по карте: %d", card.Limit))
	}
	err = e.serv.ApplicationCreate(c.Sender().Username, cardID, amount)
	if err != nil {
		log.Println(err)
		state.Set(daimyoBeginState)
		return c.Send("Возникла ошибка", keyboards.Daimyo())
	}
	err = e.serv.CardSetDisputeTrue(cardID)
	if err != nil {
		log.Println(err)
		state.Set(daimyoBeginState)
		return c.Send("Возникла ошибка", keyboards.Daimyo())
	}
	return c.Send("Данные записаны")
}

// applications
func (e *Endpoint) daimyoApplications(c tele.Context, state fsm.FSMContext) error {
	return c.Send("Выберите", keyboards.Applications())
}
func (e *Endpoint) daimyoApplicationsActive(c tele.Context, state fsm.FSMContext) error {
	applications, err := e.serv.ApplicationsGetActive()
	if err != nil {
		log.Println(err)
		return c.Send("Возникла ошибка. Попробуйте позже")
	}
	for _, application := range applications {
		c.Send(fmt.Sprintf("%s / %d", application.CardID, application.Sum))
	}
	return nil
}
func (e *Endpoint) daimyoApplicationsDisputable(c tele.Context, state fsm.FSMContext) error {
	applications, err := e.serv.ApplicationsGetDisputable()
	if err != nil {
		log.Println(err)
		return c.Send("Возникла ошибка")
	}
	state.Set(daimyoApplicationDisputableCardState)
	return c.Send("Выберите", keyboards.ApplicationList(applications))
}
func (e *Endpoint) daimyoApplicationsDisputableChoseReplenishment(c tele.Context, state fsm.FSMContext) error {
	before, _, _ := strings.Cut(c.Text(), " ")
	state.Update("cardID", before)
	log.Println(before)
	return c.Send("Выберите", keyboards.ApplicationReplenishment())
}
func (e *Endpoint) daimyoApplicationsDisputableChoseReplenishmentDone(c tele.Context, state fsm.FSMContext) error {
	cardID := state.MustGet("cardID").(string)

}
