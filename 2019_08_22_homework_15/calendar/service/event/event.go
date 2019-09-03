package event

import (
	"errors"
	"sync"
)

//go:generate protoc event.proto --go_out=plugins=grpc:. --proto_path=.

type EService struct {
	sync.RWMutex
	storage map[uint64]*Event
}

func NewService() EvInterface {
	return &EService{
		RWMutex: sync.RWMutex{},
		storage: make(map[uint64]*Event),
	}
}

func (e *EService) Create(event *Event) (*Event, error) {
	e.Lock()
	defer e.Unlock()
	id := uint64(len(e.storage))

	event.ID = id
	e.storage[id] = event

	return event, nil
}

func (e *EService) Read(eventId uint64) (*Event, error) {
	e.RLock()
	defer e.RUnlock()

	if uint64(len(e.storage)-1) < eventId {
		return nil, errors.New("event does not exists")
	}

	return e.storage[eventId], nil
}

func (e *EService) Update(event *Event) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if uint64(len(e.storage)-1) < event.ID {
		return false, errors.New("event does not exists")
	}

	e.storage[event.ID] = event

	return true, nil
}

func (e *EService) Delete(eventId uint64) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if len(e.storage) <= 0 || uint64(len(e.storage)-1) < eventId {
		return false, errors.New("event does not exists")
	}

	delete(e.storage, eventId)

	return true, nil
}
