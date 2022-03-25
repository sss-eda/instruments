package driving

import "github.com/sss-eda/instruments/internal/domain/instrument"

// CommandType TODO
type CommandType string

// Command TODO
type Command struct {
	Type         CommandType   `json:"type"`
	InstrumentID instrument.ID `json:"instrumentID"`
	Payload      []byte        `json:"payload"`
}
