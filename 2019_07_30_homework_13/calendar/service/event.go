package service

import (
	"errors"
	"sync"
)

type Event struct {
	Id          uint64
	Name        string
	Description string
	Date        string
}

type EventService struct {
	sync.RWMutex
	storage []*Event
}

func (e *EventService) Create(event *Event) (*Event, error) {
	e.Lock()
	defer e.Unlock()
	id := uint64(len(e.storage))

	event.Id = id
	e.storage = append(e.storage, event)

	return e.storage[id], nil
}

func (e *EventService) Read(eventId uint64) (*Event, error) {
	e.RLock()
	defer e.RUnlock()

	if uint64(len(e.storage)-1) < eventId {
		return nil, errors.New("event does not exists")
	}

	return e.storage[eventId], nil
}

func (e *EventService) Update(event *Event) (bool, error) {
	if uint64(len(e.storage)-1) < event.Id {
		return false, errors.New("event does not exists")
	}

	e.storage[event.Id] = event

	return true, nil
}

func (e *EventService) Delete(eventId uint64) (bool, error) {
	if len(e.storage) <= 0 || uint64(len(e.storage)-1) < eventId {
		return false, errors.New("event does not exists")
	}

	e.storage[eventId] = nil

	return true, nil
}
