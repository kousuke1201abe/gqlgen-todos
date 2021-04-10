package registry

import (
	"github.com/jinzhu/gorm"
	todoApplication "github.com/kousuke1201abe/gqlgen-todos/internal/application/todos"
	userApplication "github.com/kousuke1201abe/gqlgen-todos/internal/application/users"
	todoModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
	userModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
	"github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/database"
	todoInfrastructure "github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/todos"
	userInfrastructure "github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/users"
)

type Repository interface {
	NewTodoUsecase() todoApplication.TodoUsecase
	NewTodoRepository() todoModel.TodoRepository
	NewUserUsecase() userApplication.UserUsecase
	NewUserRepository() userModel.UserRepository
	CloseDB()
}

func NewRepository() Repository {
	return &RepositoryImpl{DB: database.NewDB()}
}

type RepositoryImpl struct {
	DB             *gorm.DB
	TodoUsecase    todoApplication.TodoUsecase
	TodoRepository todoModel.TodoRepository
	UserUsecase    userApplication.UserUsecase
	UserRepository userModel.UserRepository
}

func (r *RepositoryImpl) CloseDB() {
	database.CloseDB(r.DB)
}

func (r *RepositoryImpl) NewTodoUsecase() todoApplication.TodoUsecase {
	if r.TodoUsecase == nil {
		r.TodoUsecase = todoApplication.NewTodoUsecase(r.NewTodoRepository())
	}
	return r.TodoUsecase
}

func (r *RepositoryImpl) NewTodoRepository() todoModel.TodoRepository {
	if r.TodoRepository == nil {
		r.TodoRepository = todoInfrastructure.NewTodoRepository(r.DB)
	}
	return r.TodoRepository
}

func (r *RepositoryImpl) NewUserUsecase() userApplication.UserUsecase {
	if r.UserUsecase == nil {
		r.UserUsecase = userApplication.NewUserUsecase(r.NewUserRepository())
	}
	return r.UserUsecase
}

func (r *RepositoryImpl) NewUserRepository() userModel.UserRepository {
	if r.UserRepository == nil {
		r.UserRepository = userInfrastructure.NewUserRepository(r.DB)
	}
	return r.UserRepository
}
