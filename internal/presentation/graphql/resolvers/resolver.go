package resolvers

import "github.com/kousuke1201abe/gqlgen-todos/internal/ioc"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Registry ioc.Registry
}
