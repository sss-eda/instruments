package api

import (
	"fmt"

	"github.com/sss-eda/instruments/internal/domain/adding"
	"github.com/sss-eda/instruments/internal/domain/listing"
	"github.com/sss-eda/instruments/internal/infrastructure/memory"
	"github.com/sss-eda/instruments/internal/infrastructure/openapi"
)

// Application TODO
type Application struct {
	server Server
}

// NewApplication TODO
func NewApplication(
	serverType ServerType,
	gatewayType GatewayType,
) (*Application, error) {
	var gateway Gateway
	var err error

	switch gatewayType {
	case Memory:
		gateway, err = memory.NewGateway()
	default:
		err = fmt.Errorf("unsupported Gateway Type: %v", gatewayType)
	}
	if err != nil {
		return nil, err
	}

	var server Server

	switch serverType {
	case OpenAPI:
		server, err = openapi.NewServer(
			adding.AddInstrument(gateway),
			listing.ListInstruments(gateway),
		)
	default:
		err = fmt.Errorf("unsupport Server Type: %v", serverType)
	}
	if err != nil {
		return nil, err
	}

	return &Application{
		server: server,
	}, nil
}

// Run TODO
func (application *Application) Run() error {
	return application.server.Serve()
}
