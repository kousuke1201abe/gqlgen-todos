package repository

import "github.com/kousuke1201abe/gqlgen-todos/app/domain/model"

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
	All() ([]*model.User, error)
	FindByTodoID(todoID string) (*model.User, error)
	Find(id string) (*model.User, error)
}
