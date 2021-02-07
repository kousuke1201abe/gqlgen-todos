package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kousuke1201abe/gqlgen-todos/domain/model"
)

func (d *UserDao) TableName() string {
	return "users"
}

type UserDao struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
	All() ([]*model.User, error)
	FindByTodoID(todoID string) (*model.User, error)
	Find(id string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserDao{db: db}
}

func (d *UserDao) Create(u *model.User) (*model.User, error) {
	res := d.db.Create(u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (d *UserDao) All() ([]*model.User, error) {
	var Users []*model.User
	res := d.db.Find(&Users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return Users, nil
}

func (d *UserDao) Find(id string) (*model.User, error) {
	var Users []*model.User
	res := d.db.Where("id = ?", id).Find(&Users)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(Users) < 1 {
		return nil, nil
	}
	return Users[0], nil
}

func (d *UserDao) FindByTodoID(todoID string) (*model.User, error) {
	var users []*model.User
	res := d.db.Table("users").
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
