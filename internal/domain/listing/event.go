package listing

// EventID TODO
type EventID string

// Event TODO
type Event struct {
	ID EventID
}

// EventRepository TODO
type EventRepository interface {
	GetEventsByInstrument(instrumentID uint64, first uint64, after uint64) ([]*Event, error)
}

func ListEvents(
	repository EventRepository,
) func(uint64, uint64, uint64) ([]*Event, error) {
	return func(instrumentID uint64, first uint64, after uint64) ([]*Event, error) {
		return repository.GetEventsByInstrument(instrumentID, first, after)
	}
}
