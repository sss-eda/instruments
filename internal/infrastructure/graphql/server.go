package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// NewHTTPHandler TODO
func NewHTTPHandler() (*handler.Handler, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: RootQuery,
	})
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	}), nil
}

// PageInfo TODO
var PageInfo = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageInfo",
	Fields: graphql.Fields{
		"endCursor": &graphql.Field{
			Type: graphql.String,
		},
		"hasNextPage": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"hasPreviousPage": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"startCursor": &graphql.Field{
			Type: graphql.String,
		},
	},
})
