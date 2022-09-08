package repositories

import "gorm.io/gorm"

type Repository struct {
	User UserRepository
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
