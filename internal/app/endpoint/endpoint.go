package endpoint

import (
	"emivn/internal/app/repository"
	"emivn/internal/app/service"
	"emivn/internal/pkg/keyboards"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	tele "gopkg.in/telebot.v3"
)

var inputSG = fsm.NewStateGroup("start")

type Endpoint struct {
	serv *service.Service
	//repo *repository.Repository
}

func New(serv *service.Service, repo *repository.Repository) *Endpoint {
	return &Endpoint{
		serv: serv,
		//repo: repo,
	}
}
func (e *Endpoint) Init(bot *tele.Group, manager *fsm.Manager) {
	bot.Handle("/start", e.onStart)
	manager.Bind("/state", fsm.AnyState, func(c tele.Context, state fsm.FSMContext) error {
		return c.Send(state.State().String())
	})
	e.initSamuraiEndpoints(manager)
}

func (e *Endpoint) onStart(c tele.Context) error {
	return c.Send("Выберите", keyboards.Default())
}
