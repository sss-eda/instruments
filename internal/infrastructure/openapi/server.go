package openapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sss-eda/instruments/internal/domain/adding"
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// Server TODO
type Server struct {
	router *mux.Router
}

// Serve TODO
func (server *Server) Serve() error {
	return http.ListenAndServe(":8080", server.router)
}

// NewServer TODO
func NewServer(
	addInstrument adding.AddInstrumentUseCase,
	listInstruments listing.ListInstrumentsUseCase,
) (*Server, error) {
	router := mux.NewRouter()

	// GET https://api.sansa.dev/instruments
	router.Handle("/instruments", listInstrumentsHandlerFunc(listInstruments)).Methods("GET")
	// POST https://api.sansa.dev/instruments
	router.Handle("/instruments", addInstrumentHandlerFunc(addInstrument)).Methods("POST")

	return &Server{router: router}, nil
}

func listInstrumentsHandlerFunc(
	listInstruments listing.ListInstrumentsUseCase,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func addInstrumentHandlerFunc(
	addInstrument adding.AddInstrumentUseCase,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
