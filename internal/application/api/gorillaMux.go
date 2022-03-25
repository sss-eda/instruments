package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// NewGorillaMuxHTTPRouter TODO
func NewGorillaMuxHTTPRouter(
	instruments listing.InstrumentRepository,
) (*mux.Router, error) {
	router := mux.NewRouter()

	router.HandleFunc(
		"/instruments",
		ListInstruments(listing.Instruments(instruments)),
	)

	return router, nil
}

// ListInstruments TODO
func ListInstruments(
	instrumentLister listing.ListInstrumentsUseCase,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
