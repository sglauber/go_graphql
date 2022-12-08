package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sglauber/studiosol/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("StudioSol - GraphiQL", "/query"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/graphql for StudioSol - GraphiQL", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
