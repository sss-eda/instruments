package graphql

import "github.com/graphql-go/graphql"

var (
	// Command TODO
	Command *graphql.Object
	// InstrumentCommandEdge TODO
	InstrumentCommandEdge *graphql.Object
	// InstrumentCommandConnection TODO
	InstrumentCommandConnection *graphql.Object
)

func init() {
	Command = graphql.NewObject(graphql.ObjectConfig{
		Name: "Command",
		Interfaces: []*graphql.Interface{
			NodeInterface,
		},
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Name: "",
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
	})

	InstrumentCommandEdge = graphql.NewObject(graphql.ObjectConfig{
		Name: "InstrumentCommandEdge",
		Interfaces: []*graphql.Interface{
			EdgeInterface,
		},
		Fields: graphql.Fields{
			"cursor": &graphql.Field{
				Name: "",
				Type: graphql.NewNonNull(graphql.String),
			},
			"node": &graphql.Field{
				Name: "",
				Type: Command,
			},
		},
	})

	InstrumentCommandConnection = graphql.NewObject(graphql.ObjectConfig{
		Name: "InstrumentCommandConnection",
		Interfaces: []*graphql.Interface{
			ConnectionInterface,
		},
		Fields: graphql.Fields{
			"edges": &graphql.Field{
				Name:        "",
				Type:        graphql.NewList(InstrumentCommandEdge),
				Description: "List of instrument-command edges in connection",
			},
			"nodes": &graphql.Field{
				Name:        "",
				Type:        graphql.NewList(Command),
				Description: "List of command nodes in connection",
			},
			"pageInfo": &graphql.Field{
				Name: "",
				Type: graphql.NewNonNull(PageInfo),
			},
			"totalCount": &graphql.Field{
				Name: "",
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
	})
}
