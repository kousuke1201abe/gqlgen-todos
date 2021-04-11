package dto

import (
	"github.com/kousuke1201abe/gqlgen-todos/internal/domain/models/users"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ConvertUser(user *users.User) (*User, error) {
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func ConvertUsers(users []*users.User) ([]*User, error) {
	var results []*User
	for _, user := range users {
		results = append(results, &User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}
