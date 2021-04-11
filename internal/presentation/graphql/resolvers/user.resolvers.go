package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	todoModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
	userModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
)

func (r *userResolver) Todos(ctx context.Context, obj *userModel.User) ([]*todoModel.Todo, error) {
	return r.TodoApplicationService.FindByUserID(obj.ID)
}

// User returns generated1.UserResolver implementation.
func (r *Resolver) User() generated1.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
