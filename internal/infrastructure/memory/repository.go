package memory

import (
	"sort"

	"github.com/sss-eda/instruments/internal/domain/adding"
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// Repository TODO
type Repository struct {
	instruments []*Instrument
}

// NewGateway TODO
func NewGateway() (*Repository, error) {
	return &Repository{
		instruments: make([]*Instrument, 0),
	}, nil
}

// AddInstrument TODO
func (repository *Repository) AddInstrument(
	id uint64,
	instrument *adding.Instrument,
) error {
	index := sort.Search(len(repository.instruments), func(i int) bool {
		return repository.instruments[i].ID > InstrumentID(id)
	})

	repository.instruments = append(repository.instruments, &Instrument{})
	copy(repository.instruments[index+1:], repository.instruments[index:])
	repository.instruments[index] = &Instrument{
		ID:   InstrumentID(id),
		Name: instrument.Name,
	}

	return nil
}

// GetInstruments TODO
func (repository *Repository) GetInstruments(
	first uint64,
	after uint64,
) ([]*listing.Instrument, error) {
	index := sort.Search(len(repository.instruments), func(i int) bool {
		return repository.instruments[i].ID > InstrumentID(after)
	})

	var instruments []*listing.Instrument
	for i, instrument := range repository.instruments[index:] {
		if uint64(i) > first {
			break
		}
		instruments = append(instruments, &listing.Instrument{
			ID:   uint64(instrument.ID),
			Name: instrument.Name,
		})
	}

	return instruments, nil
}
