package repository

import (
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/domain"
)

type EvInterface interface {
	Create(*domain.Event) (*domain.Event, error)
	Read(eventId uint64) (*domain.Event, error)
	Update(*domain.Event) (bool, error)
	Delete(eventId uint64) (bool, error)
}
