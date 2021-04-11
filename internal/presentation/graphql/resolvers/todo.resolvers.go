package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/dto"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
)

func (r *todoResolver) User(ctx context.Context, obj *dto.Todo) (*dto.User, error) {
	user, _ := r.UserApplicationService.FindByTodoID(obj.ID)
	return &dto.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Todo returns generated1.TodoResolver implementation.
func (r *Resolver) Todo() generated1.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
