package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	dto1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/dto"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
)

func (r *mutationResolver) UpdateTodo(ctx context.Context, input dto1.UpdateTodo) (*dto1.Todo, error) {
	todo, _ := r.TodoApplicationService.Update(input.TodoID, input.Text)
	return dto1.ConvertTodo(todo)
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input dto1.NewTodo) (*dto1.Todo, error) {
	todo, _ := r.TodoApplicationService.Create(input.Text, input.UserID)
	return dto1.ConvertTodo(todo)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input dto1.DeleteTodo) (*dto1.Todo, error) {
	todo, _ := r.TodoApplicationService.Delete(input.TodoID)
	return dto1.ConvertTodo(todo)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input dto1.NewUser) (*dto1.User, error) {
	user, _ := r.UserApplicationService.Create(input.Name)
	return dto1.ConvertUser(user)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
