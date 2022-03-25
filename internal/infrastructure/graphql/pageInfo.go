package graphql

import "github.com/graphql-go/graphql"

// PageInfo TODO
var PageInfo = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageInfo",
	Fields: graphql.Fields{
		"endCursor": &graphql.Field{
			Name: "",
			Type: graphql.String,
		},
		"hasNextPage": &graphql.Field{
			Name: "",
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"hasPreviousPage": &graphql.Field{
			Name: "",
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"startCursor": &graphql.Field{
			Name: "",
			Type: graphql.String,
		},
	},
})
