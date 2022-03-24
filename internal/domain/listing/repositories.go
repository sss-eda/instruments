package listing

import "context"

// InstrumentRepository TODO
type InstrumentRepository interface {
	List(context.Context) ([]*Instrument, error)
	GetByID(context.Context, InstrumentID) (*Instrument, error)
}
