package graphql

import "github.com/graphql-go/graphql"

var (
	// Instrument TODO
	Instrument *graphql.Object
	// InstrumentEdge TODO
	InstrumentEdge *graphql.Object
	// InstrumentConnection TODO
	InstrumentConnection *graphql.Object
)

func init() {
	Instrument = graphql.NewObject(graphql.ObjectConfig{
		Name: "Instrument",
		Interfaces: []*graphql.Interface{
			NodeInterface,
		},
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
		},
	})

	InstrumentEdge = graphql.NewObject(graphql.ObjectConfig{
		Name: "InstrumentEdge",
		Interfaces: []*
		Fields: graphql.Fields{
			"cursor": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"node": &graphql.Field{
				Type: Node,
			},
		},
	})

	InstrumentConnection = graphql.NewObject(graphql.ObjectConfig{
		Name: "InstrumentConnection",
		Fields: graphql.Fields{
			"edges": InstrumentEdge,
			"pageInfo": &graphql.Field{
				Type: PageInfo,
			},
		},
	})
}
