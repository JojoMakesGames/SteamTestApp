package main

import (
	"log"
	"net/http"
	"os"
	"server/graph"
	"server/helpers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	driver, err := neo4j.NewDriver(os.Getenv("NEO4J_URI"), neo4j.BasicAuth("", "", ""))
	if err != nil {
		print(err)
	}

	resolver := graph.Resolver{GameService: helpers.GameService{Driver: driver}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
