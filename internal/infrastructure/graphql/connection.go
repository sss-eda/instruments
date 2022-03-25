package graphql

import "github.com/graphql-go/graphql"

var (
	// ConnectionInterface TODO
	ConnectionInterface *graphql.Interface
)

func init() {
	ConnectionInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "Connection",
		Fields: graphql.Fields{
			"edges": &graphql.Field{
				Name: "",
				Type: graphql.NewList(InstrumentEdge),
				// Type:        graphql.NewList(EdgeInterface),
				Description: "List of edges in connection",
			},
			"nodes": &graphql.Field{
				Name: "",
				Type: graphql.NewList(Instrument),
				// Type:        graphql.NewList(NodeInterface),
				Description: "List of nodes in connection",
			},
			"pageInfo": &graphql.Field{
				Name:        "",
				Type:        graphql.NewNonNull(PageInfo),
				Description: "PageInfo",
			},
			"totalCount": &graphql.Field{
				Name:        "",
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Total Count",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			return InstrumentConnection
		},
		Description: "Connection Interface",
	})
}
