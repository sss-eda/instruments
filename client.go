package instruments

import "github.com/sss-eda/instruments/internal/domain/listing"

// Client TODO
type Client struct {
	app Application
}

// Service TODO
type Service interface {

}

// NewClient TODO
func NewClient(
	clientType ClientType,
) *Client {
	client := &Client{}
	switch clientType {
	case OpenAPI:
		client.listingService = openapi.
	case GraphQL:
	case GRPC:
	}
}


instruments.NewClient(instruments.OpenAPI)