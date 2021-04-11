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
		newTodoApplicationService,
		newUserApplicationService,
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

func newResolver(todoApplicationService todoApplication.TodoApplicationService, userApplicationService userApplication.UserApplicationService, db *gorm.DB) *resolvers.Resolver {
	return &resolvers.Resolver{
		DB:                     db,
		TodoApplicationService: todoApplicationService,
		UserApplicationService: userApplicationService,
	}
}

func newTodoApplicationService(repo todoModel.TodoRepository) todoApplication.TodoApplicationService {
	return todoApplication.NewTodoApplicationService(repo)
}

func newTodoRepository(db *gorm.DB) todoModel.TodoRepository {
	return todoInfrastructure.NewTodoRepository(db)
}

func newUserApplicationService(repo userModel.UserRepository) userApplication.UserApplicationService {
	return userApplication.NewUserApplicationService(repo)
}

func newUserRepository(db *gorm.DB) userModel.UserRepository {
	return userInfrastructure.NewUserRepository(db)
}
