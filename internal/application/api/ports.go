package api

import "github.com/sss-eda/instruments/internal/domain/listing"

// listingService - This is the port, along with the DTOs
type listingService interface {
	Instruments(listing.InstrumentsInput) (listing.InstrumentsOutput, error)
	Commands(listing.CommandsInput) (listing.CommandsOutput, error)
	Events(listing.EventsInput) (listing.EventsOutput, errors)
}
