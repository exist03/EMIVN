package endpoint

import (
	"emivn/internal/models"
	"emivn/internal/pkg/keyboards"
	"fmt"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
	"log"
)

var (
	daimyoBeginState           = inputSG.New("daimyoBeginState")
	daimyoHierarchyState       = inputSG.New("daimyoHierarchyState")
	daimyoSamuraiNicknameState = inputSG.New("daimyoSamuraiNicknameState")
	daimyoSamuraiUsernameState = inputSG.New("daimyoSamuraiUsernameState")
)

func (e *Endpoint) initDaimyoEndpoints(manager *fsm.Manager) {
	manager.Bind(&keyboards.BtnDaimyo, fsm.DefaultState, e.daimyo)
	manager.Bind(&keyboards.BtnHierarchy, daimyoBeginState, e.daimyoHierarchy)

	manager.Bind(&keyboards.BtnSubordinates, daimyoHierarchyState, e.daimyoSubordinates)
	manager.Bind(&keyboards.BtnCreateSamurai, daimyoHierarchyState, e.daimyoCreateSamurai)

	manager.Bind(tele.OnText, daimyoSamuraiNicknameState, e.daimyoCreateSamuraiNickname)
	manager.Bind(tele.OnText, daimyoSamuraiUsernameState, e.daimyoCreateSamuraiUsername)
}
func (e *Endpoint) daimyo(c tele.Context, state fsm.FSMContext) error {
	if !e.serv.Repo.ValidUser(c.Sender().Username, "Daimyo") {
		return c.Send("У вас нет прав")
	}
	state.Set(daimyoBeginState)
	return c.Send("Выберите действие", keyboards.Daimyo())
}

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
