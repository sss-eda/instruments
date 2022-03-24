package graphql

import (
	"github.com/graphql-go/graphql"
)

// Node TODO
type Node struct {
	ID string
}

// NodeInterface TODO
var NodeInterface = graphql.NewInterface(graphql.InterfaceConfig{
	Name:        "Node",
	Description: "A node on a connection edge",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.ID),
			Description: "The node's unique ID.",
		},
	},
	ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
		return Instrument
	},
})
