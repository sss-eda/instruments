package openapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// Server TODO
type Server struct{}

// NewServer TODO
func NewServer(
	listService listing.Service,
) (*Server, error) {
	router := mux.NewRouter()

	// GET https://api.sansa.dev/instruments
	router.HandleFunc("/", getInstruments(listService)).Methods("GET")

	return server, nil
}

func getInstruments(
	service listing.Service,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instruments, err := service.ListInstruments(r.Context())
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, err)
			return
		}

		response, err := json.Marshal(instruments)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, err)
			return
		}

		_, err = w.Write(response)
		if err != nil {
			log.Println(err)
		}
	}
}
