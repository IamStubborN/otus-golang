package repository

import (
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/config"
)

func NewStorage(cfg *config.Config) (EvInterface, error) {
	var storage EvInterface
	var err error

	switch cfg.Storage.Provider {
	case "postgres":
		storage, err = NewDatabase(cfg)
	case "cache":
		storage = NewCache()
	}

	return storage, err
}
