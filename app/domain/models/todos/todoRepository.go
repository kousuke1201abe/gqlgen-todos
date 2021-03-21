package todos

type TodoRepository interface {
	Create(u *Todo) (*Todo, error)
	All() ([]*Todo, error)
	FindByUserID(userID string) ([]*Todo, error)
	Find(id string) (*Todo, error)
	Destroy(t *Todo) error
	Save(t *Todo) error
}
