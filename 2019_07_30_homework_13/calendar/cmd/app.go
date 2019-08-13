package cmd

import (
	"github.com/IamStubborN/otus-golang/2019_07_30_homework_13/calendar/config"
	"github.com/IamStubborN/otus-golang/2019_07_30_homework_13/calendar/logger"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger *logrus.Logger
}

func NewApp() *App {
	app := &App{}
	cfg := initializeConfig()
	app.initializeLogger(cfg)
	return app
}

func initializeConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalln(err)
	}
	return cfg
}

func (app *App) initializeLogger(cfg *config.Config) {
	log, err := logger.NewLogger(cfg.Logger)
	if err != nil {
		app.Logger.Fatalln(err)
	}
	app.Logger = log
}

func (app *App) Run() {
	app.Logger.Info("App run")
	//TODO
}
