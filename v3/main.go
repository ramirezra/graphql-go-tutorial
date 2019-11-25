package main

import (
	"context"
	"encoding/json"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

const schemaString = `
schema {
	query: Query
}

type Query {
	greet: String!
	greetPerson(person: String!): String!
	greetPersonTimeOfDay(person: String!, timeOfDay: TimeOfDay!): String!
}

# Enumerate times of day:
enum TimeOfDay{
	MORNING
	AFTERNOON
	EVENING
}
`

// RootResolver type defined
type RootResolver struct{}

// Greet function defined
func (*RootResolver) Greet() string {
	return "Hello, world!"
}

// GreetPerson defined
func (*RootResolver) GreetPerson(args struct{ Person string }) string {
	return fmt.Sprintf("Hello, %s!", args.Person)
}

// PersonTimeOfDaysArgs type defined
type PersonTimeOfDaysArgs struct {
	Person    string
	TimeOfDay string
}

// TimesOfDay map defined
var TimesOfDay = map[string]string{
	"MORNING":   "Good morning",
	"AFTERNOON": "Good afternoon",
	"EVENING":   "Good evening",
}

// GreetPersonTimeOfDay function defined
func (*RootResolver) GreetPersonTimeOfDay(ctx context.Context, args PersonTimeOfDaysArgs) string {
	timeOfDay, ok := TimesOfDay[args.TimeOfDay]
	if !ok {
		timeOfDay = "Go to bed"
	}
	return fmt.Sprintf("%s, %s!", timeOfDay, args.Person)
}

// Schema defined
var Schema = graphql.MustParseSchema(schemaString, &RootResolver{})

func main() {
	ctx := context.Background()

	type ClientQuery struct {
		OpName    string
		Query     string
		Variables map[string]interface{}
	}

	q1 := ClientQuery{
		OpName: "Greet",
		Query: `query Greet{
			greet }`,
		Variables: nil,
	}
	resp1 := Schema.Exec(ctx, q1.Query, q1.OpName, q1.Variables)

	json1, err := json.MarshalIndent(resp1, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json1))

	q2 := ClientQuery{
		OpName: "GreetPerson",
		Query: `query GreetPerson($person: String!){
			greetPerson(person: $person)
			}`,
		Variables: map[string]interface{}{
			"person": "Robinson",
		},
	}

	resp2 := Schema.Exec(ctx, q2.Query, q2.OpName, q2.Variables)
	json2, err := json.MarshalIndent(resp2, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json2))

	q3 := ClientQuery{
		OpName: "GreetPersonTimeOfDay",
		Query: `query GreetPersonTimeOfDay($person: String!, $timeOfDay: TimeOfDay!){
			greetPersonTimeOfDay(person: $person, timeOfDay: $timeOfDay)
			}`,
		Variables: map[string]interface{}{
			"person":    "Robinson",
			"timeOfDay": "MORNING",
		},
	}

	resp3 := Schema.Exec(ctx, q3.Query, q3.OpName, q3.Variables)
	json3, err := json.MarshalIndent(resp3, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json3))
}
