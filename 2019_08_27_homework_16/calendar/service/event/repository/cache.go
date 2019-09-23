package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/domain"
)

type EventsCache struct {
	sync.RWMutex
	storage map[uint64]*domain.Event
}

func NewCache() EvInterface {
	return &EventsCache{
		RWMutex: sync.RWMutex{},
		storage: make(map[uint64]*domain.Event),
	}
}

func (e *EventsCache) Create(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	e.Lock()
	defer e.Unlock()
	id := uint64(len(e.storage))

	event.ID = id
	e.storage[id] = event

	return event, nil
}

func (e *EventsCache) Read(ctx context.Context, eventId uint64) (*domain.Event, error) {
	e.RLock()
	defer e.RUnlock()

	if uint64(len(e.storage)-1) < eventId {
		return nil, errors.New("event does not exists")
	}

	return e.storage[eventId], nil
}

func (e *EventsCache) Update(ctx context.Context, event *domain.Event) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if uint64(len(e.storage)-1) < event.ID {
		return false, errors.New("event does not exists")
	}

	e.storage[event.ID] = event

	return true, nil
}

func (e *EventsCache) Delete(ctx context.Context, eventId uint64) (bool, error) {
	e.Lock()
	defer e.Unlock()

	if len(e.storage) <= 0 || uint64(len(e.storage)-1) < eventId {
		return false, errors.New("event does not exists")
	}

	delete(e.storage, eventId)

	return true, nil
}
