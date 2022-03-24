package graphql

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/sss-eda/instruments/internal/driving"
)

// Client TODO
type Client struct {
	router *mux.Router
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
	return nil
}

// GetCommands TODO
func (client *Client) GetCommands(
	ctx context.Context,
) error {
	return nil
}
