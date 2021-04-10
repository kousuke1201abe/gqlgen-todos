package ioc

import (
	"github.com/jinzhu/gorm"
	todoApplication "github.com/kousuke1201abe/gqlgen-todos/internal/application/todos"
	userApplication "github.com/kousuke1201abe/gqlgen-todos/internal/application/users"
	todoModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/todos"
	userModel "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
	"github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/database"
	todoInfrastructure "github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/todos"
	userInfrastructure "github.com/kousuke1201abe/gqlgen-todos/internal/infrastructure/users"
	"github.com/kousuke1201abe/gqlgen-todos/internal/presentation/graphql/resolvers"
	dig "go.uber.org/dig"
)

func Dig() *dig.Container {
	c := dig.New()

	registerDependencies(
		c,
		newDB,
		newResolver,
		newTodoUsecase,
		newUserUsecase,
		newTodoRepository,
		newUserRepository,
	)

	return c
}

func registerDependencies(c *dig.Container, constructors ...interface{}) {
	for i := 0; i < len(constructors); i++ {
		err := c.Provide(constructors[i])
		if err != nil {
			panic(err)
		}
	}
}

func newDB() *gorm.DB {
	return database.NewDB()
}

func newResolver(todoUsecase todoApplication.TodoUsecase, userUsecase userApplication.UserUsecase, db *gorm.DB) *resolvers.Resolver {
	return &resolvers.Resolver{
		DB:          db,
		TodoUsecase: todoUsecase,
		UserUsecase: userUsecase,
	}
}

func newTodoUsecase(repo todoModel.TodoRepository) todoApplication.TodoUsecase {
	return todoApplication.NewTodoUsecase(repo)
}

func newTodoRepository(db *gorm.DB) todoModel.TodoRepository {
	return todoInfrastructure.NewTodoRepository(db)
}

func newUserUsecase(repo userModel.UserRepository) userApplication.UserUsecase {
	return userApplication.NewUserUsecase(repo)
}

func newUserRepository(db *gorm.DB) userModel.UserRepository {
	return userInfrastructure.NewUserRepository(db)
}
