package openapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// NewServer TODO
func NewServer(
	listService listing.Service,
) (http.Handler, error) {
	router := mux.NewRouter()

	// https://api.sansa.dev/instruments
	router.Handle("/instruments", Instruments(listService))

	return router, nil
}
