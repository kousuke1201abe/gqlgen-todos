package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	model "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models"
	"github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
	"github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/graphql/generated"
)

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*todos.Todo, error) {
	return r.Repository.NewTodoUsecase().Update(input.TodoID, input.Text)
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*todos.Todo, error) {
	return r.Repository.NewTodoUsecase().Create(input.Text, input.UserID)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*todos.Todo, error) {
	return r.Repository.NewTodoUsecase().Delete(input.TodoID)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*users.User, error) {
	return r.Repository.NewUserUsecase().Create(input.Name)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
