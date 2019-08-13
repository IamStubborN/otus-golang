package service

import (
	"errors"
	"sync"

	"github.com/IamStubborN/otus-golang/2019_07_30_homework_13/calendar/models"
)

//go:generate protoc event.proto --go_out=$PWD/models --proto_path=$PWD/models

type EventService struct {
	sync.RWMutex
	storage map[uint64]*models.Event
}

func (e *EventService) Create(event *models.Event) (*models.Event, error) {
	e.Lock()
	defer e.Unlock()
	id := uint64(len(e.storage))

	event.ID = id
	e.storage[id] = event

	return event, nil
}

func (e *EventService) Read(eventId uint64) (*models.Event, error) {
	e.RLock()
	defer e.RUnlock()

	if uint64(len(e.storage)-1) < eventId {
		return nil, errors.New("event does not exists")
	}

	return e.storage[eventId], nil
}

func (e *EventService) Update(event *models.Event) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if uint64(len(e.storage)-1) < event.ID {
		return false, errors.New("event does not exists")
	}

	e.storage[event.ID] = event

	return true, nil
}

func (e *EventService) Delete(eventId uint64) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if len(e.storage) <= 0 || uint64(len(e.storage)-1) < eventId {
		return false, errors.New("event does not exists")
	}

	delete(e.storage, eventId)

	return true, nil
}
