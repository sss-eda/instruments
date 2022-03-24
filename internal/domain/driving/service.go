package driving

import "context"

// Service TODO
type Service interface {
	SendCommand(context.Context, *Command) error
}
