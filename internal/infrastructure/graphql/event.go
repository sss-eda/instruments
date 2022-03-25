package graphql

import "github.com/graphql-go/graphql"

// var (
// 	// Event TODO
// 	Event *graphql.Object
// 	// InstrumentEventEdge TODO
// 	InstrumentEventEdge *graphql.Object
// 	// InstrumentEventConnection TODO
// 	InstrumentEventConnection *graphql.Object
// )

// func init() {

// Event TODO
var Event = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
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

// InstrumentEventEdge TODO
var InstrumentEventEdge = graphql.NewObject(graphql.ObjectConfig{
	Name: "InstrumentEventEdge",
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
			Type: Event,
		},
	},
})

// InstrumentEventConnection TODO
var InstrumentEventConnection = graphql.NewObject(graphql.ObjectConfig{
	Name: "InstrumentEventConnection",
	Interfaces: []*graphql.Interface{
		ConnectionInterface,
	},
	Fields: graphql.Fields{
		"edges": &graphql.Field{
			Name:        "",
			Type:        graphql.NewList(InstrumentEventEdge),
			Description: "List of instrument-event edges in connection",
		},
		"nodes": &graphql.Field{
			Name:        "",
			Type:        graphql.NewList(Event),
			Description: "List of event nodes in connection",
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

// }
