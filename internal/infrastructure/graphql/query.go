package graphql

import "github.com/graphql-go/graphql"

// Query TODO
var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"instruments": &graphql.Field{
				Type: InstrumentConnection,
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 10,
						Description:  "",
					},
					"after": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
						Description:  "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
		Description: "The Root Query",
	},
)
