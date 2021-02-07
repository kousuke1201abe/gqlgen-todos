package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kousuke1201abe/gqlgen-todos/app/domain/model"
	"github.com/kousuke1201abe/gqlgen-todos/graphql/generated"
)

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	return r.Repository.NewTodoUsecase().Update(input.TodoID, input.Text)
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return r.Repository.NewTodoUsecase().Create(input.Text, input.UserID)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
	return r.Repository.NewTodoUsecase().Delete(input.TodoID)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Repository.NewUserUsecase().Create(input.Name)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
