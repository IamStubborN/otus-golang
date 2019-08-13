package logger

import (
	"io/ioutil"
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
	if IsValid(config.Output) {
		logFile, err := os.Create(config.Output)
		if err != nil {
			return nil, errors.Wrap(err, "can't create log file")
		}
		logger.SetOutput(logFile)
	}

	return logger, nil
}

func IsValid(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}

	return false
}
