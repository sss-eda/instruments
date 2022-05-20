package api

import "fmt"

// Server TODO
type Server interface {
	Serve() error
}

// ServerType TODO
type ServerType int

const (
	// OpenAPI TODO
	OpenAPI ServerType = iota
	// GraphQL TODO
	GraphQL
	// GRPC TODO
	GRPC
	// NATS TODO
	NATS
)

// MarshalText TODO
func (serverType *ServerType) MarshalText() (text []byte, err error) {
	switch *serverType {
	case OpenAPI:
		text = []byte("OpenAPI")
	case GraphQL:
		text = []byte("GraphQL")
	case GRPC:
		text = []byte("GRPC")
	case NATS:
		text = []byte("NATS")
	default:
		text = []byte("Undefined")
		err = fmt.Errorf("could not marshal ClientType=%d to text", serverType)
	}

	return
}

// UnmarshalText TODO
func (serverType *ServerType) UnmarshalText(text []byte) (err error) {
	switch string(text) {
	case "OpenAPI":
		*serverType = OpenAPI
	case "GraphQL":
		*serverType = GraphQL
	case "GRPC":
		*serverType = GRPC
	case "NATS":
		*serverType = NATS
	default:
		*serverType = -1
		err = fmt.Errorf("could not unmarshal ClientType from text=\"%q\"", string(text))
	}

	return
}

// String TODO
func (serverType ServerType) String() string {
	textBytes, err := serverType.MarshalText()
	text := string(textBytes)
	if err != nil {
		return fmt.Sprintf("%q: %q", text, err.Error())
	}

	return text
}
