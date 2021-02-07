package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kousuke1201abe/gqlgen-todos/database"
	"github.com/kousuke1201abe/gqlgen-todos/domain/model"
	"github.com/kousuke1201abe/gqlgen-todos/graph/generated"
)

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	log.Printf("[mutationResolver.UpdateTodo] input: %#v", input)
	todoRepository := database.NewTodoRepository(r.DB)
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
	todo, err := database.NewTodoRepository(r.DB).Create(&model.Todo{
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
	todoRepository := database.NewTodoRepository(r.DB)
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
	user, err := database.NewUserRepository(r.DB).Create(&model.User{
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

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	log.Printf("[queryResolver.Todo]")
	todos, err := database.NewTodoRepository(r.DB).All()
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
	todo, err := database.NewTodoRepository(r.DB).Find(id)
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
	users, err := database.NewUserRepository(r.DB).All()
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
	user, err := database.NewUserRepository(r.DB).Find(id)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	log.Printf("[todoResolver.User]")
	user, err := database.NewUserRepository(r.DB).FindByTodoID(obj.ID)
	if err != nil {
		return nil, err
	}
	result := &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
	return result, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	log.Println("[userResolver.Todos]")
	todos, err := database.NewTodoRepository(r.DB).FindByUserID(obj.ID)
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
