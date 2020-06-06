package schemas

import (
	"community-api/repositories"
	"community-api/services/topic"

	"github.com/graphql-go/graphql"
)

var topicType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Topic",
		Fields: graphql.Fields{
			"code": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var topicRepository = topic.NewTopicInMemoryRepository()

// New returns a new Community API GraphQL Schema
func New() (graphql.Schema, error) {
	fields := graphql.Fields{
		"topics": &graphql.Field{
			Type:        graphql.NewList(topicType),
			Description: "Get topic list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return topicRepository.List(), nil
			},
		},
		"topic": &graphql.Field{
			Type:        topicType,
			Description: "Get a topic by code",
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				codeString, ok := p.Args["code"].(string)
				if ok {
					code := repositories.TopicCode(codeString)
					return topicRepository.Get(code)
				}
				return nil, nil //TODO: return error
			},
		},
	}
	queryType := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: fields})

	return graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
}
