package api

import (
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// ListInstruments TODO
func ListInstruments(
	list listingService,
) func(func(*ListInstrumentsResponse) error, *ListInstrumentsRequest) {
	return func(
		respond func(*ListInstrumentsResponse) error,
		request *ListInstrumentsRequest,
	) {
		instrumentsOutput, err := list.Instruments(
			&listing.InstrumentsInput{
				First: request.First,
				After: request.After,
			},
		)
		if err != nil {
			respond(&ListInstrumentsResponse{
				Instruments: nil,
				Error:       err,
			})
			return
		}

		instruments := make([]*Instrument, len(instrumentsOutput))
		for i, instrument := range instruments {
			instruments[i] = &Instrument{
				ID: instrument.ID,
			}
		}

		respond(&ListInstrumentsResponse{
			Instruments: []*Instrument{},
			Error:       nil,
		})
	}
}

// ListInstrumentsRequest TODO
type ListInstrumentsRequest struct {
	First int
	After string
}

// ListInstrumentsResponse TODO
type ListInstrumentsResponse struct {
	Instruments []*Instrument
	Error       error
}
