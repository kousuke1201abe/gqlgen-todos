package database

import (
	"github.com/jinzhu/gorm"
	"github.com/kousuke1201abe/gqlgen-todos/domain/model"
)

func (d *TodoDao) TableName() string {
	return "todos"
}

type TodoDao struct {
	db *gorm.DB
}

type TodoRepository interface {
	Create(u *model.Todo) (*model.Todo, error)
	All() ([]*model.Todo, error)
	FindByUserID(userID string) ([]*model.Todo, error)
	Find(id string) (*model.Todo, error)
	Destroy(t *model.Todo) error
	Save(t *model.Todo) error
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoDao{db: db}
}

func (d *TodoDao) Create(t *model.Todo) (*model.Todo, error) {
	res := d.db.Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (d *TodoDao) All() ([]*model.Todo, error) {
	var todos []*model.Todo
	res := d.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *TodoDao) Find(id string) (*model.Todo, error) {
	var todos []*model.Todo
	res := d.db.Where("id = ?", id).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(todos) < 1 {
		return nil, nil
	}
	return todos[0], nil
}

func (d *TodoDao) FindByUserID(userID string) ([]*model.Todo, error) {
	var todos []*model.Todo
	res := d.db.Where("user_id = ?", userID).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *TodoDao) Destroy(t *model.Todo) error {
	res := d.db.Delete(t)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (d *TodoDao) Save(t *model.Todo) error {
	res := d.db.Save(t)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
