package database

import "github.com/jinzhu/gorm"

type Todo struct {
	ID     string `gorm:"column:id;primary_key"`
	Text   string `gorm:"column:text"`
	Done   bool   `gorm:"column:done"`
	UserID string `gorm:"column:user_id"`
}

func (u *Todo) TableName() string {
	return "todos"
}

type TodoDao struct {
	db *gorm.DB
}

type TodoRepository interface {
	Create(u *Todo) (*Todo, error)
	All() ([]*Todo, error)
	FindByUserID(userID string) ([]*Todo, error)
	Find(id string) (*Todo, error)
	Destroy(t *Todo) (*Todo, error)
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoDao{db: db}
}

func (d *TodoDao) Create(t *Todo) (*Todo, error) {
	res := d.db.Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (d *TodoDao) All() ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *TodoDao) Find(id string) (*Todo, error) {
	var todos []*Todo
	res := d.db.Where("id = ?", id).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(todos) < 1 {
		return nil, nil
	}
	return todos[0], nil
}

func (d *TodoDao) FindByUserID(userID string) ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Where("user_id = ?", userID).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *TodoDao) Destroy(t *Todo) (*Todo, error) {
	res := d.db.Delete(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	return t, nil
}
