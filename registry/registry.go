package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/kousuke1201abe/gqlgen-todos/config/database"
	"github.com/kousuke1201abe/gqlgen-todos/domain/repository"
	"github.com/kousuke1201abe/gqlgen-todos/infrastructure"
	"github.com/kousuke1201abe/gqlgen-todos/usecase"
)

type Repository interface {
	NewTodoUsecase() usecase.TodoUsecase
	NewTodoRepository() repository.TodoRepository
	NewUserUsecase() usecase.UserUsecase
	NewUserRepository() repository.UserRepository
	CloseDB()
}

func NewRepository() Repository {
	return &RepositoryImpl{DB: database.NewDB()}
}

type RepositoryImpl struct {
	DB             *gorm.DB
	TodoUsecase    usecase.TodoUsecase
	TodoRepository repository.TodoRepository
	UserUsecase    usecase.UserUsecase
	UserRepository repository.UserRepository
}

func (r *RepositoryImpl) CloseDB() {
	database.CloseDB(r.DB)
	return
}

func (r *RepositoryImpl) NewTodoUsecase() usecase.TodoUsecase {
	if r.TodoUsecase == nil {
		r.TodoUsecase = usecase.NewTodoUsecase(r.NewTodoRepository())
	}
	return r.TodoUsecase
}

func (r *RepositoryImpl) NewTodoRepository() repository.TodoRepository {
	if r.TodoRepository == nil {
		r.TodoRepository = infrastructure.NewTodoRepository(r.DB)
	}
	return r.TodoRepository
}

func (r *RepositoryImpl) NewUserUsecase() usecase.UserUsecase {
	if r.UserUsecase == nil {
		r.UserUsecase = usecase.NewUserUsecase(r.NewUserRepository())
	}
	return r.UserUsecase
}

func (r *RepositoryImpl) NewUserRepository() repository.UserRepository {
	if r.UserRepository == nil {
		r.UserRepository = infrastructure.NewUserRepository(r.DB)
	}
	return r.UserRepository
}
