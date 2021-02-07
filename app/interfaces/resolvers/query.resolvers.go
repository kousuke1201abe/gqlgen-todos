package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kousuke1201abe/gqlgen-todos/app/domain/model"
	"github.com/kousuke1201abe/gqlgen-todos/graphql/generated"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.Repository.NewTodoUsecase().All()
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	return r.Repository.NewTodoUsecase().Find(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.Repository.NewUserUsecase().All()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.Repository.NewUserUsecase().Find(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
