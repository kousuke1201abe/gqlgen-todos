package resolvers

import (
	"github.com/jinzhu/gorm"
	"github.com/kousuke1201abe/gqlgen-todos/internal/application/todos"
	"github.com/kousuke1201abe/gqlgen-todos/internal/application/users"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB                     *gorm.DB
	TodoApplicationService todos.TodoApplicationService
	UserApplicationService users.UserApplicationService
}
