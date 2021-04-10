package todos

import (
	"github.com/google/uuid"
	model "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
)

type TodoUsecase interface {
	Create(text string, userID string) (*model.Todo, error)
	Update(id string, text string) (*model.Todo, error)
	Delete(id string) (*model.Todo, error)
	All() ([]*model.Todo, error)
	Find(id string) (*model.Todo, error)
	FindByUserID(id string) ([]*model.Todo, error)
}

type TodoUsecaseImpl struct {
	TodoRepository model.TodoRepository
}

func NewTodoUsecase(TodoRepository model.TodoRepository) TodoUsecase {
	return &TodoUsecaseImpl{TodoRepository: TodoRepository}
}

func (tu *TodoUsecaseImpl) Create(text string, userID string) (*model.Todo, error) {
	id, _ := uuid.NewRandom()
	todo, err := tu.TodoRepository.Create(&model.Todo{
		ID:     id.String(),
		Text:   text,
		Done:   false,
		UserID: userID,
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

func (tu *TodoUsecaseImpl) Update(id string, text string) (*model.Todo, error) {
	todoRepository := tu.TodoRepository
	todo, err := todoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	todo.Text = text
	err = todoRepository.Save(todo)
	if err != nil {
		return nil, err
	}
	todo, err = todoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (tu *TodoUsecaseImpl) Delete(id string) (*model.Todo, error) {
	todoRepository := tu.TodoRepository
	todo, err := todoRepository.Find(id)
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

func (tu *TodoUsecaseImpl) All() ([]*model.Todo, error) {
	todos, err := tu.TodoRepository.All()
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

func (tu *TodoUsecaseImpl) Find(id string) (*model.Todo, error) {
	todo, err := tu.TodoRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (tu *TodoUsecaseImpl) FindByUserID(id string) ([]*model.Todo, error) {
	todos, err := tu.TodoRepository.FindByUserID(id)
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
