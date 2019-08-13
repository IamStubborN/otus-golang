package logger

import (
	"os"

	"github.com/IamStubborN/otus-golang/2019_07_30_homework_13/calendar/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewLogger(config config.Logger) (*logrus.Logger, error) {
	logger := logrus.New()
	lvl, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse logger level")
	}

	logger.SetLevel(lvl)
	logger.SetOutput(os.Stdout)

	return logger, nil
}
