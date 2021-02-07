package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/kousuke1201abe/gqlgen-todos/domain/model"
	"github.com/kousuke1201abe/gqlgen-todos/graph/generated"
)

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	log.Println("[userResolver.Todos]")
	todos, err := r.Repository.NewTodoRepository().FindByUserID(obj.ID)
	if err != nil {
		return nil, err
	}
	var results []*model.Todo
	for _, todo := range todos {
		results = append(results, &model.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
