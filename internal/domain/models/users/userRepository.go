package users

type UserRepository interface {
	Create(u *User) (*User, error)
	All() ([]*User, error)
	FindByTodoID(todoID string) (*User, error)
	Find(id string) (*User, error)
}
