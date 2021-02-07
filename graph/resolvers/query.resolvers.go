package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/kousuke1201abe/gqlgen-todos/domain/model"
	"github.com/kousuke1201abe/gqlgen-todos/graph/generated"
	"github.com/kousuke1201abe/gqlgen-todos/infrastructure"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	log.Printf("[queryResolver.Todo]")
	todos, err := infrastructure.NewTodoRepository(r.DB).All()
	var results []*model.Todo
	if err != nil {
		return nil, err
	}
	for _, todo := range todos {
		results = append(results, &model.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	log.Printf("[queryResolver.User] id: %s", id)
	todo, err := infrastructure.NewTodoRepository(r.DB).Find(id)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	log.Printf("[queryResolver.Users]")
	users, err := infrastructure.NewUserRepository(r.DB).All()
	var results []*model.User
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		results = append(results, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	log.Printf("[queryResolver.User] id: %s", id)
	user, err := infrastructure.NewUserRepository(r.DB).Find(id)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
