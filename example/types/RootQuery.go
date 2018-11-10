package types

import (
	"github.com/afandylamusu/go-graphiql/example/logic"
	"github.com/graphql-go/graphql"
)

// RootQuery ...
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getMessage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				msg := logic.GetMessage()
				return msg, nil
			},
		},
	},
})
