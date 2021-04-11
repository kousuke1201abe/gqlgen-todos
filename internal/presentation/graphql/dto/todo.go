package dto

import (
	"github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
)

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user_id"`
}

func ConvertTodo(todo *todos.Todo) (*Todo, error) {
	return &Todo{
		ID:     todo.ID,
		Text:   todo.Text,
		Done:   todo.Done,
		UserID: todo.UserID,
	}, nil
}

func ConvertTodos(todos []*todos.Todo) ([]*Todo, error) {
	var results []*Todo
	for _, todo := range todos {
		results = append(results, &Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
		})
	}
	return results, nil
}
