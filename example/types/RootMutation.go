package types

import (
	"github.com/afandylamusu/go-graphiql/example/logic"
	"github.com/graphql-go/graphql"
)

// RootMutation ...
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"setMessage": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"msg": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				msg := p.Args["msg"].(string)
				logic.SetMessage(msg)
				return msg, nil
			},
		},
	},
})
