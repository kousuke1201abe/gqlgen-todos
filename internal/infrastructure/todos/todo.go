package todos

import (
	"github.com/jinzhu/gorm"
	model "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
)

func (t *TodoRepository) TableName() string {
	return "todos"
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) model.TodoRepository {
	return &TodoRepository{db: db}
}

func (t *TodoRepository) Create(m *model.Todo) (*model.Todo, error) {
	res := t.db.Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (t *TodoRepository) All() ([]*model.Todo, error) {
	var todos []*model.Todo
	res := t.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodoRepository) Find(id string) (*model.Todo, error) {
	var todos []*model.Todo
	res := t.db.Where("it = ?", id).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(todos) < 1 {
		return nil, nil
	}
	return todos[0], nil
}

func (t *TodoRepository) FindByUserID(userId string) ([]*model.Todo, error) {
	var todos []*model.Todo
	res := t.db.Where("user_it = ?", userId).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodoRepository) Destroy(m *model.Todo) error {
	res := t.db.Delete(m)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (t *TodoRepository) Save(m *model.Todo) error {
	res := t.db.Save(m)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
