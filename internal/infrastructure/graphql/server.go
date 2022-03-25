package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// NewHTTPHandler TODO
func NewHTTPHandler() (*handler.Handler, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: Query,
		// Types: []graphql.Type{
		// 	NodeInterface,
		// 	EdgeInterface,
		// 	ConnectionInterface,
		// },
	})
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	}), nil
}
