package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	todoModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
	userModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
	generated1 "github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/generated"
)

func (r *todoResolver) User(ctx context.Context, obj *todoModel.Todo) (*userModel.User, error) {
	log.Printf("[todoResolver.User]")
	userRepo := r.Repository.NewUserRepository()
	user, err := userRepo.FindByTodoID(obj.ID)
	if err != nil {
		return nil, err
	}
	result := &userModel.User{
		ID:   user.ID,
		Name: user.Name,
	}
	return result, nil
}

// Todo returns generated1.TodoResolver implementation.
func (r *Resolver) Todo() generated1.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
