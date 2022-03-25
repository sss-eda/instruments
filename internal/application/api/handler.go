package api

import (
	"github.com/gorilla/mux"
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// NewGorillaMuxRouter TODO
func NewGorillaMuxRouter(
	instruments listing.InstrumentRepository,
) (*mux.Router, error) {
	router := mux.NewRouter()

	router.HandleFunc(
		"/instruments",
		ListInstruments(listing.Instruments(instruments)),
	)

	return router, nil
}
