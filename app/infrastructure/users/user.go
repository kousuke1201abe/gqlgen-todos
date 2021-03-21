package users

import (
	"github.com/jinzhu/gorm"
	model "github.com/kousuke1201abe/gqlgen-todos/app/domain/models/users"
)

func (u *UserRepository) TableName() string {
	return "users"
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(m *model.User) (*model.User, error) {
	res := u.db.Create(m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (u *UserRepository) All() ([]*model.User, error) {
	var Users []*model.User
	res := u.db.Find(&Users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return Users, nil
}

func (u *UserRepository) Find(id string) (*model.User, error) {
	var Users []*model.User
	res := u.db.Where("id = ?", id).Find(&Users)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(Users) < 1 {
		return nil, nil
	}
	return Users[0], nil
}

func (u *UserRepository) FindByTodoID(todoID string) (*model.User, error) {
	var users []*model.User
	res := u.db.Table("users").
		Select("users.*").
		Joins("LEFT JOIN todos ON todos.user_id = users.id").
		Where("todos.id = ?", todoID).
		First(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	if users == nil || len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
}
