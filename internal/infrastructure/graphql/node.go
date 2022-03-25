package graphql

import (
	"github.com/graphql-go/graphql"
)

var (
	// NodeInterface TODO
	NodeInterface *graphql.Interface
)

func init() {
	NodeInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Node",
		Description: "A node on a connection edge",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Name:        "",
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "The node's unique ID.",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			return Instrument
		},
	})
}
