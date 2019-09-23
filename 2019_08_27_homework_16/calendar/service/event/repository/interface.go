package repository

import (
	"context"
	"github.com/IamStubborN/otus-golang/2019_08_27_homework_16/calendar/service/event/domain"
)

type EvInterface interface {
	Create(ctx context.Context, ev *domain.Event) (*domain.Event, error)
	Read(ctx context.Context, eventId uint64) (*domain.Event, error)
	Update(ctx context.Context, ev *domain.Event) (bool, error)
	Delete(ctx context.Context, eventId uint64) (bool, error)
}
