package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/database"
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
	dig := ioc.Dig()

	err := dig.Invoke(func(r *resolvers.Resolver) error {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		defer database.CloseDB(r.DB)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
		return nil
	})

	if err != nil {
		panic(err)
	}
}
