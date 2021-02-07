package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/kousuke1201abe/gqlgen-todos/config/database"
	"github.com/kousuke1201abe/gqlgen-todos/domain/repository"
	"github.com/kousuke1201abe/gqlgen-todos/infrastructure"
)

type Repository interface {
	NewUserRepository() repository.UserRepository
	NewTodoRepository() repository.TodoRepository
	CloseDB()
}

func NewRepository() Repository {
	return &RepositoryImpl{DB: database.NewDB()}
}

type RepositoryImpl struct {
	DB             *gorm.DB
	TodoRepository repository.TodoRepository
	UserRepository repository.UserRepository
}

func (r *RepositoryImpl) CloseDB() {
	database.CloseDB(r.DB)
	return
}

func (r *RepositoryImpl) NewTodoRepository() repository.TodoRepository {
	if r.TodoRepository == nil {
		r.TodoRepository = infrastructure.NewTodoRepository(r.DB)
	}
	return r.TodoRepository
}

func (r *RepositoryImpl) NewUserRepository() repository.UserRepository {
	if r.UserRepository == nil {
		r.UserRepository = infrastructure.NewUserRepository(r.DB)
	}
	return r.UserRepository
}
