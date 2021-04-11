package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/dto"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
)

func (r *userResolver) Todos(ctx context.Context, obj *dto.User) ([]*dto.Todo, error) {
	todos, _ := r.TodoApplicationService.FindByUserID(obj.ID)
	var results []*dto.Todo
	for _, todo := range todos {
		results = append(results, &dto.Todo{
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
