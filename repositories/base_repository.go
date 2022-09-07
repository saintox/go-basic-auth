package repositories

type Repository struct {
	User UserRepository
}

func NewRepository() *Repository {
	return &Repository{
		User: NewUserRepository(),
	}
}
