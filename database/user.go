package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID   string `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (u *User) TableName() string {
	return "users"
}

type UserDao struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(u *User) (*User, error)
	All() ([]*User, error)
	FindByTodoID(todoID string) (*User, error)
	Find(id string) (*User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserDao{db: db}
}

func (d *UserDao) Create(u *User) (*User, error) {
	res := d.db.Create(u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (d *UserDao) All() ([]*User, error) {
	var Users []*User
	res := d.db.Find(&Users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return Users, nil
}

func (d *UserDao) Find(id string) (*User, error) {
	var Users []*User
	res := d.db.Where("id = ?", id).Find(&Users)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(Users) < 1 {
		return nil, nil
	}
	return Users[0], nil
}

func (d *UserDao) FindByTodoID(todoID string) (*User, error) {
	var users []*User
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
