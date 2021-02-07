package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kousuke1201abe/gqlgen-todos/domain/model"
	"github.com/kousuke1201abe/gqlgen-todos/graph/generated"
	"github.com/kousuke1201abe/gqlgen-todos/infrastructure"
)

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	log.Printf("[mutationResolver.UpdateTodo] input: %#v", input)
	todoRepository := infrastructure.NewTodoRepository(r.DB)
	todo, err := todoRepository.Find(input.TodoID)
	if err != nil {
		return nil, err
	}
	todo.Text = input.Text
	err = todoRepository.Save(todo)
	if err != nil {
		return nil, err
	}
	todo, err = todoRepository.Find(input.TodoID)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	log.Printf("[mutationResolver.CreateTodo] input: %#v", input)
	id, _ := uuid.NewRandom()
	todo, err := infrastructure.NewTodoRepository(r.DB).Create(&model.Todo{
		ID:     id.String(),
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
	log.Printf("[mutationResolver.DeleteTodo] input: %#v", input)
	todoRepository := infrastructure.NewTodoRepository(r.DB)
	todo, err := todoRepository.Find(input.TodoID)
	if err != nil {
		return nil, err
	}
	err = todoRepository.Destroy(todo)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Printf("[mutationResolver.CreateUser] input: %#v", input)
	id, _ := uuid.NewRandom()
	user, err := infrastructure.NewUserRepository(r.DB).Create(&model.User{
		ID:   id.String(),
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
