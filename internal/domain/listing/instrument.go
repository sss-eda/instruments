package listing

// Instrument TODO
type Instrument struct {
	ID       uint64
	Name     string
	Commands []*Command
	Events   []*Event
}

// InstrumentRepository TODO
type InstrumentRepository interface {
	GetInstruments(first uint64, after uint64) ([]*Instrument, error)
}

func ListInstruments(
	repository InstrumentRepository,
) ListInstrumentsUseCase {
	return func(first uint64, after uint64) ([]*Instrument, error) {
		return repository.GetInstruments(first, after)
	}
}

type ListInstrumentsUseCase func(uint64, uint64) ([]*Instrument, error)
