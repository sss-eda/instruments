package graphql

import "github.com/graphql-go/graphql"

// EdgeInterface TODO
var EdgeInterface = graphql.NewInterface(graphql.InterfaceConfig{
	Name: "Edge",
	Fields: graphql.Fields{
		// 	"node": &graphql.Field{
		// 		Name: "",
		// 		Type: Instrument,
		// 		// Type:        NodeInterface,
		// 		Description: "Node",
		// 	},
		"cursor": &graphql.Field{
			Name:        "",
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Cursor",
		},
	},
	// ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
	// 	return InstrumentEdge
	// },
	Description: "An edge in a connection",
})

func init() {
	EdgeInterface.AddFieldConfig(
		"node",
		&graphql.Field{
			Name:        "",
			Type:        NodeInterface,
			Description: "Node",
		},
	)

	// EdgeInterface.AddFieldConfig(
	// 	"cursor",
	// 	&graphql.Field{
	// 		Name:        "",
	// 		Type:        graphql.NewNonNull(graphql.String),
	// 		Description: "Cursor",
	// 	},
	// )

	EdgeInterface.ResolveType = func(p graphql.ResolveTypeParams) *graphql.Object {
		return InstrumentEdge
	}
}
