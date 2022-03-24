package openapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/sss-eda/instruments/internal/domain/listing"
)

// Instruments TODO
func Instruments(
	repository listing.InstrumentRepository,
) http.Handler {
	list := listing.Instruments(repository)

	handler := &instrumentsHandler{
		service: listingService,
	}

	return handler
}

// instrumentsHandler TODO
type instrumentsHandler struct {
	service listing.Service
}

// HandleHTTP TODO
func (handler *instrumentsHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

}

// ListInstrumentsHandlerFunc TODO
func ListInstrumentsHandlerFunc(
	list,
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

// ListInstruments TODO - This code belongs in the same area as where the actual
// server is run. We need: the route, the method
func ListInstruments() error {
	resp, err := http.Post(
		"http://localhost:8080/instruments",
		"application/json",
		strings.NewReader(`{"hello": "world"}`),
	)
	if err != nil {
		return err
	}
	log.Println(resp)

	return nil
}

// // Client TODO
// type Client struct {
// 	router *mux.Router
// 	driver driving.Service
// }

// // NewClient TODO
// func NewClient() (*Client, error) {
// 	return &Client{}, nil
// }

// // SendCommand TODO
// func (client *Client) SendCommand(
// 	ctx context.Context,
// 	command driving.Command,
// ) error {
// 	resp, err := http.Post(
// 		"http://localhost:8080/instruments",
// 		"application/json",
// 		strings.NewReader(`{"hello": "world"}`),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println(resp)

// 	return nil
// }

// // GetCommands TODO
// func (client *Client) GetCommands(
// 	ctx context.Context,
// ) error {
// 	return nil
// }
