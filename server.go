package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kousuke1201abe/gqlgen-todos/graph/generated"
	"github.com/kousuke1201abe/gqlgen-todos/interfaces/resolvers"
	"github.com/kousuke1201abe/gqlgen-todos/registry"
)

func main() {
	const defaultPort = "5050"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	repositry := registry.NewRepository()
	defer repositry.CloseDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{Repository: repositry}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
