package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/kousuke1201abe/gqlgen-todos/domain/model"
	"github.com/kousuke1201abe/gqlgen-todos/graph/generated"
)

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	log.Printf("[todoResolver.User]")
	user, err := r.Repository.NewUserRepository().FindByTodoID(obj.ID)
	if err != nil {
		return nil, err
	}
	result := &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
	return result, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
