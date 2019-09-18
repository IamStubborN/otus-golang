package event

type EvInterface interface {
	Create(*Event) (*Event, error)
	Read(eventId uint64) (*Event, error)
	Update(*Event) (bool, error)
	Delete(eventId uint64) (bool, error)
}
