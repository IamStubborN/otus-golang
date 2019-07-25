package logger

import (
	"github.com/IamStubborN/otus-golang/2019_07_18_homework_12/microservice/cmd/config"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	SetupLogger()
}

func GetLogger() *logrus.Logger {
	return logger
}

func SetupLogger() {
	logger = logrus.New()
	logger.SetLevel(logrus.Level(config.GetConfig().LogLevel))
	logger.Infof("Logger initialized with %d level", logger.Level)
}
