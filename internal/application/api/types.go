package api

import "fmt"

// Type TODO
type Type int

const (
	// OpenAPI TODO
	OpenAPI Type = iota
	// GraphQL TODO
	GraphQL
	// GRPC TODO
	GRPC
	// NATS TODO
	NATS
)

// MarshalText TODO
func (t *Type) MarshalText() (text []byte, err error) {
	switch *t {
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
		err = fmt.Errorf("could not marshal ClientType=%d to text", t)
	}

	return
}

// UnmarshalText TODO
func (t *Type) UnmarshalText(text []byte) (err error) {
	switch string(text) {
	case "OpenAPI":
		*t = OpenAPI
	case "GraphQL":
		*t = GraphQL
	case "GRPC":
		*t = GRPC
	case "NATS":
		*t = NATS
	default:
		*t = -1
		err = fmt.Errorf("could not unmarshal ClientType from text=\"%q\"", string(text))
	}

	return
}

// String TODO
func (t Type) String() string {
	textBytes, err := t.MarshalText()
	text := string(textBytes)
	if err != nil {
		return fmt.Sprintf("%q: %q", text, err.Error())
	}

	return text
}
