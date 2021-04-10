package users

import (
	"github.com/google/uuid"
	model "github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
)

type UserUsecase interface {
	Create(name string) (*model.User, error)
	All() ([]*model.User, error)
	Find(id string) (*model.User, error)
	FindByTodoID(id string) (*model.User, error)
}

type UserUsecaseImpl struct {
	UserRepository model.UserRepository
}

func NewUserUsecase(UserRepository model.UserRepository) UserUsecase {
	return &UserUsecaseImpl{UserRepository: UserRepository}
}

func (uu *UserUsecaseImpl) Create(name string) (*model.User, error) {
	id, _ := uuid.NewRandom()
	user, err := uu.UserRepository.Create(&model.User{
		ID:   id.String(),
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (uu *UserUsecaseImpl) All() ([]*model.User, error) {
	users, err := uu.UserRepository.All()
	var results []*model.User
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		results = append(results, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}

func (uu *UserUsecaseImpl) Find(id string) (*model.User, error) {
	user, err := uu.UserRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (uu *UserUsecaseImpl) FindByTodoID(id string) (*model.User, error) {
	user, err := uu.UserRepository.FindByTodoID(id)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
