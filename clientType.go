package instruments

import (
	"fmt"
)

// ClientType TODo
type ClientType int

const (
	// OpenAPI TODO
	OpenAPI ClientType = iota
	// GraphQL TODO
	GraphQL
	// GRPC TODO
	GRPC
)

// MarshalText TODO
func (clientType *ClientType) MarshalText() (text []byte, err error) {
	switch *clientType {
	case OpenAPI:
		text = []byte("OpenAPI")
	case GraphQL:
		text = []byte("GraphQL")
	case GRPC:
		text = []byte("GRPC")
	default:
		text = []byte("Undefined")
		err = fmt.Errorf("could not marshal ClientType=%d to text", clientType)
	}

	return
}

// UnmarshalText TODO
func (clientType *ClientType) UnmarshalText(text []byte) (err error) {
	switch string(text) {
	case "OpenAPI":
		*clientType = OpenAPI
	case "GraphQL":
		*clientType = GraphQL
	case "GRPC":
		*clientType = GRPC
	default:
		*clientType = -1
		err = fmt.Errorf("could not unmarshal ClientType from text=\"%q\"", string(text))
	}

	return
}

// String TODO
func (clientType ClientType) String() string {
	textBytes, err := clientType.MarshalText()
	text := string(textBytes)
	if err != nil {
		return fmt.Sprintf("%q: %q", text, err.Error())
	}

	return text
}
