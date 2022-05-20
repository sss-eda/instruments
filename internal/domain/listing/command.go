package listing

// CommandID TODO
type CommandID string

// Command TODO
type Command struct {
	ID CommandID
}

// CommandRepository TODO
type CommandRepository interface {
	GetCommandsByInstrument(instrumentID uint64, first uint64, after uint64) ([]*Command, error)
}

func ListCommands(
	repository CommandRepository,
) func(uint64, uint64, uint64) ([]*Command, error) {
	return func(instrumentID uint64, first uint64, after uint64) ([]*Command, error) {
		return repository.GetCommandsByInstrument(instrumentID, first, after)
	}
}
