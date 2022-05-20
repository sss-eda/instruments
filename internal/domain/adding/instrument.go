package adding

import "math/rand"

// Instrument TODO
type Instrument struct {
	Name string
}

// InstrumentRepository TODO
type InstrumentRepository interface {
	AddInstrument(uint64, *Instrument) error
}

// AddInstrument TODO
func AddInstrument(
	repository InstrumentRepository,
) AddInstrumentUseCase {
	return func(instrument *Instrument) error {
		id := rand.Uint64()
		return repository.AddInstrument(id, instrument)
	}
}

type AddInstrumentUseCase func(*Instrument) error
