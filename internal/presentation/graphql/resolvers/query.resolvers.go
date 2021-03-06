package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	dto1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/dto"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*dto1.Todo, error) {
	todos, _ := r.TodoApplicationService.All()
	return dto1.ConvertTodos(todos)
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*dto1.Todo, error) {
	todo, _ := r.TodoApplicationService.Find(id)
	return dto1.ConvertTodo(todo)
}

func (r *queryResolver) Users(ctx context.Context) ([]*dto1.User, error) {
	users, _ := r.UserApplicationService.All()
	return dto1.ConvertUsers(users)
}

func (r *queryResolver) User(ctx context.Context, id string) (*dto1.User, error) {
	user, _ := r.UserApplicationService.Find(id)
	return dto1.ConvertUser(user)
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
