package listing

// Instrument TODO
type Instrument struct {
	id InstrumentID
}

// ID TODO
func (instrument *Instrument) ID() InstrumentID {
	return instrument.id
}

// InstrumentID TODO
type InstrumentID string
