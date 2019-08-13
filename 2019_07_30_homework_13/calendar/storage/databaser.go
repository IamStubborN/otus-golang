package storage

import (
	"github.com/IamStubborN/otus-golang/2019_07_30_homework_13/calendar/models"
)

type Databaser interface {
	Create(*models.Event) (*models.Event, error)
	Read(eventId uint64) (*models.Event, error)
	Update(*models.Event) (bool, error)
	Delete(eventId uint64) (bool, error)
}
