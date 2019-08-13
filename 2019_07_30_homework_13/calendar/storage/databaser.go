package storage

import "github.com/IamStubborN/otus-golang/2019_07_30_homework_13/calendar/service"

type Databaser interface {
	Create(*service.Event) (*service.Event, error)
	Read(eventId uint64) (*service.Event, error)
	Update(*service.Event) (bool, error)
	Delete(eventId uint64) (bool, error)
}
