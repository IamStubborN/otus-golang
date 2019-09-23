package app

import (
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/config"
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/logger"
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger   *logrus.Logger
	EService *event.Service
}

func NewApp() *App {
	app := &App{}
	cfg := initializeConfig()
	app.Logger = initializeLogger(cfg)
	app.EService = initializeEventService(cfg, app.Logger)

	return app
}

func initializeConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalln(err)
	}

	return cfg
}

func initializeEventService(cfg *config.Config, logger *logrus.Logger) *event.Service {
	EService, err := event.NewEventService(cfg)
	if err != nil {
		logger.Fatalln(err)
	}

	return EService
}

func initializeLogger(cfg *config.Config) *logrus.Logger {
	log, err := logger.NewLogger(cfg.Logger)
	if err != nil {
		logrus.Fatalln(err)
	}

	return log
}

func (app *App) Run() {
	app.Logger.Info("App run")

	go app.EService.Server.Run(app.Logger)
	app.EService.Client.Run(app.Logger)

	app.Logger.Info("App exit")
}
