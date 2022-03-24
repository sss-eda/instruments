package openapi

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sss-eda/instruments/internal/driving"
)

// Client TODO
type Client struct {
	router *mux.Router
	driver driving.Service
}

// NewClient TODO
func NewClient() (*Client, error) {
	return &Client{}, nil
}

// SendCommand TODO
func (client *Client) SendCommand(
	ctx context.Context,
	command driving.Command,
) error {
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

// GetCommands TODO
func (client *Client) GetCommands(
	ctx context.Context,
) error {
	return nil
}
