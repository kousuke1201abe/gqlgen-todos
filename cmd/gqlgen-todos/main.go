package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kousuke1201abe/gqlgen-todos/internal/ioc"
	"github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
	"github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/resolvers"
)

func main() {
	const defaultPort = "5050"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	registry := ioc.NewRegistry()
	defer registry.CloseDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{Registry: registry}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
