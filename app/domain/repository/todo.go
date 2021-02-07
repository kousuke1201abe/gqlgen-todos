package repository

import "github.com/kousuke1201abe/gqlgen-todos/app/domain/model"

type TodoRepository interface {
	Create(u *model.Todo) (*model.Todo, error)
	All() ([]*model.Todo, error)
	FindByUserID(userID string) ([]*model.Todo, error)
	Find(id string) (*model.Todo, error)
	Destroy(t *model.Todo) error
	Save(t *model.Todo) error
}
