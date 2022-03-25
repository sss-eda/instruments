package graphql

import "github.com/graphql-go/graphql"

var (
	// Query TODO
	Query *graphql.Object
)

func init() {
	Query = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"instruments": &graphql.Field{
				Type: InstrumentConnection,
				// Args: graphql.FieldConfigArgument{
				// 	"first": &graphql.ArgumentConfig{
				// 		Type:         graphql.Int,
				// 		DefaultValue: 10,
				// 		Description:  "",
				// 	},
				// 	"after": &graphql.ArgumentConfig{
				// 		Type:         graphql.String,
				// 		DefaultValue: "",
				// 		Description:  "",
				// 	},
				// },
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return instrumentConnection, nil
				},
			},
		},
		Description: "The Root Query",
	})
}
