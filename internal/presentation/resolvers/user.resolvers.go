package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	todoModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
	userModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/graphql/generated"
)

func (r *userResolver) Todos(ctx context.Context, obj *userModel.User) ([]*todoModel.Todo, error) {
	log.Println("[userResolver.Todos]")
	todoRepo := r.Repository.NewTodoRepository()
	todos, err := todoRepo.FindByUserID(obj.ID)
	if err != nil {
		return nil, err
	}
	var results []*todoModel.Todo
	for _, todo := range todos {
		results = append(results, &todoModel.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}

// User returns generated1.UserResolver implementation.
func (r *Resolver) User() generated1.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
