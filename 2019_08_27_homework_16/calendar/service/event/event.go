package event

import (
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/config"
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/delivery/gc"
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/repository"
)

type Service struct {
	Client *gc.Client
	Server gc.Server
}

func NewEventService(cfg *config.Config) (*Service, error) {
	storage, err := repository.NewStorage(cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		Client: gc.NewClient(),
		Server: gc.NewServer(storage),
	}, nil
}
