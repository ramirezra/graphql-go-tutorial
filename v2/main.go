package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	graphql "github.com/graph-gophers/graphql-go"
)

const schemaString = `
	# Define the schema
	schema {
		query: Query
	}

	# Define the queries for the schema
	type Query {
		greet: String!
	}
`

// RootResolver defined
type RootResolver struct{}

// Greet function defined
func (*RootResolver) Greet() string {
	return "Hello, world!"
}

// There are two way to define the schema:
//  graphql.MustParseSechem(...) *graphql.Schema // Panics on error.
//  graphql.ParseSchema(...) (*graphql.Schema, error) // Let's you handle the error

var Schema = graphql.MustParseSchema(schemaString, &RootResolver{})

func main() {
	query := `query Greet{
		greet
		}`

	ctx := context.Background()

	resp := Schema.Exec(ctx, query, "", nil)
	json, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(json))
}
