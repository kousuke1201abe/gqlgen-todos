package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
	"github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*todos.Todo, error) {
	return r.Repository.NewTodoUsecase().All()
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*todos.Todo, error) {
	return r.Repository.NewTodoUsecase().Find(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*users.User, error) {
	return r.Repository.NewUserUsecase().All()
}

func (r *queryResolver) User(ctx context.Context, id string) (*users.User, error) {
	return r.Repository.NewUserUsecase().Find(id)
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
