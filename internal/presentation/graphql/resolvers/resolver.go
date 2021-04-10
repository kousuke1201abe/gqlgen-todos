package resolvers

//go:generate go run github.com/99designs/gqlgen

import "github.com/kousuke1201abe/gqlgen-todos/internal/registry"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository registry.Repository
}
