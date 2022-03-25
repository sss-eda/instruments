package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// var (
// 	// Instrument TODO
// 	Instrument *graphql.Object
// 	// InstrumentEdge TODO
// 	InstrumentEdge *graphql.Object
// 	// InstrumentConnection TODO
// 	InstrumentConnection *graphql.Object
// )

// func init() {

// Instrument TODO
var Instrument = graphql.NewObject(graphql.ObjectConfig{
	Name: "Instrument",
	Interfaces: []*graphql.Interface{
		NodeInterface,
	},
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if instrument, ok := p.Source.(SomeInstrument); ok {
					return instrument.ID, nil
				}
				return nil, fmt.Errorf("something went wrong in instrument object")
			},
			Description: "Instrument's unique ID.",
		},
		// "commands": &graphql.Field{
		// 	Type: InstrumentCommandConnection,
		// },
		// "events": &graphql.Field{
		// 	Type: InstrumentEventConnection,
		// },
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		return true
	},
	Description: "An Instrument.",
})

// InstrumentEdge TODO
var InstrumentEdge = graphql.NewObject(graphql.ObjectConfig{
	Name: "InstrumentEdge",
	Interfaces: []*graphql.Interface{
		EdgeInterface,
	},
	Fields: graphql.Fields{
		"cursor": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if edge, ok := p.Source.(SomeInstrumentEdge); ok {
					return edge.Cursor, nil
				}
				return nil, fmt.Errorf("InstrumentEdge: cursor")
			},
		},
		"node": &graphql.Field{
			Type: Instrument,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if edge, ok := p.Source.(SomeInstrumentEdge); ok {
					return edge.Node, nil
				}
				return nil, fmt.Errorf("InstrumentEdge: node")
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		return true
	},
	Description: "Edge in instrument connections.",
})

// InstrumentConnection TODO
var InstrumentConnection = graphql.NewObject(graphql.ObjectConfig{
	Name: "InstrumentConnection",
	Interfaces: []*graphql.Interface{
		ConnectionInterface,
	},
	Fields: graphql.Fields{
		"edges": &graphql.Field{
			Type: graphql.NewList(InstrumentEdge),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if connection, ok := p.Source.(SomeInstrumentConnection); ok {
					return connection.Edges, nil
				}
				return nil, fmt.Errorf("InstrumentConnection: edges")
			},
			Description: "List of instrument edges in connection",
		},
		"nodes": &graphql.Field{
			Type: graphql.NewList(Instrument),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if connection, ok := p.Source.(SomeInstrumentConnection); ok {
					return connection.Nodes, nil
				}
				return nil, fmt.Errorf("InstrumentConnection: nodes")
			},
			Description: "List of instrument nodes in connection",
		},
		"pageInfo": &graphql.Field{
			Type: graphql.NewNonNull(PageInfo),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if connection, ok := p.Source.(SomeInstrumentConnection); ok {
					return connection.PageInfo, nil
				}
				return nil, fmt.Errorf("InstrumentConnection: pageInfo")
			},
		},
		"totalCount": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if connection, ok := p.Source.(SomeInstrumentConnection); ok {
					return connection.TotalCount, nil
				}
				return nil, fmt.Errorf("InstrumentConnection: totalCount")
			},
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		return true
	},
})

// }
