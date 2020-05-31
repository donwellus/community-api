package schemas

import (
	"github.com/graphql-go/graphql"
)

// Topic contains information about one topic
type Topic struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

var sampleTopics = []Topic{
	{
		Code: "xpto",
		Name: "XPTO",
	},
	{
		Code: "otpx",
		Name: "OTPX",
	},
	{
		Code: "internet_speed",
		Name: "Internet Speed",
	},
}

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

func getTopicList() []Topic {
	return sampleTopics
}

// New returns a new Community API GraphQL Schema
func New() (graphql.Schema, error) {
	fields := graphql.Fields{
		"topics": &graphql.Field{
			Type:        graphql.NewList(topicType),
			Description: "Get topic list",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return getTopicList(), nil
			},
		},
	}
	queryType := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: fields})

	return graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
}
