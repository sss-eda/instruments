package graphql

import (
	"github.com/graphql-go/graphql"
)

// var (
// 	// NodeInterface TODO
// 	NodeInterface *graphql.Interface
// )

// func init() {

// NodeInterface TODO
var NodeInterface = graphql.NewInterface(graphql.InterfaceConfig{
	Name: "Node",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.ID),
			Description: "The node's unique ID.",
		},
	},
	// ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
	// 	return Instrument
	// },
	Description: "A node on a connection edge",
})

func init() {
	NodeInterface.ResolveType = func(p graphql.ResolveTypeParams) *graphql.Object {
		return Instrument
	}
}

// }
