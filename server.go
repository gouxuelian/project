package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
)

var db DB

func init() {
	var err error
	d, err := newDB("./employees.sqlite3")
	if err != nil {
		panic(err)
	}

	db = *d
}

// nolint: gas
func main() {
	s, err := getSchema("./schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s, &Resolver{})

	http.Handle("/graphiql", staticFile("graphiql.html"))
	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe("localhost:9080", nil))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
func staticFile(filename string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filename)
	})
}
