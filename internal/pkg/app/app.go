package app

import (
	"emivn/internal/app/endpoint"
	"emivn/internal/app/repository"
	"emivn/internal/app/service"
	"emivn/pkg/db"
	"github.com/joho/godotenv"
	fsm "github.com/vitaliy-ukiru/fsm-telebot"
	"github.com/vitaliy-ukiru/fsm-telebot/storages/memory"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"os"
	"time"
)

type App struct {
	handler    *endpoint.Endpoint
	serv       *service.Service
	repository *repository.Repository
}

func New() (*App, error) {
	db, err := db.OpenDB("quest:quest@/EMIVN")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	app := &App{}
	app.repository = repository.New(db)
	app.serv = service.New(app.repository)
	app.handler = endpoint.New(app.serv, app.repository)
	return app, nil
}

func (a *App) Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	bot, err := tele.NewBot(tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
	}
	bot.Use(middleware.AutoRespond())
	bGroup := bot.Group()
	storage := memory.NewStorage()
	defer storage.Close()
	manager := fsm.NewManager(bGroup, memory.NewStorage())
	a.handler.Init(bGroup, manager)
	log.Println("Handlers configured")
	bot.Start()
	return nil
}
