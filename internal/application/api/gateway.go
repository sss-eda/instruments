package api

import (
	"fmt"

	"github.com/sss-eda/instruments/internal/domain/adding"
	"github.com/sss-eda/instruments/internal/domain/listing"
)

// Gateway TODO
type Gateway interface {
	adding.InstrumentRepository
	listing.InstrumentRepository
}

// GatewayType TODO
type GatewayType int

const (
	// Memory TODO
	Memory GatewayType = iota
)

// MarshalText TODO
func (gatewayType *GatewayType) MarshalText() (text []byte, err error) {
	switch *gatewayType {
	case Memory:
		text = []byte("Memory")
	default:
		text = []byte("Undefined")
		err = fmt.Errorf("could not marshal GatewayType=%d to text", gatewayType)
	}

	return
}

// UnmarshalText TODO
func (gatewayType *GatewayType) UnmarshalText(text []byte) (err error) {
	switch string(text) {
	case "OpenAPI":
		*gatewayType = Memory
	default:
		*gatewayType = -1
		err = fmt.Errorf("could not unmarshal GatewayType from text=\"%q\"", string(text))
	}

	return
}

// String TODO
func (gatewayType GatewayType) String() string {
	textBytes, err := gatewayType.MarshalText()
	text := string(textBytes)
	if err != nil {
		return fmt.Sprintf("%q: %q", text, err.Error())
	}

	return text
}
