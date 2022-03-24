package listing

import "context"

// Service TODO
type Service interface {
	ListInstruments(context.Context) ([]*Instrument, error)
}
