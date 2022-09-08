package repositories

import (
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, ID string)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) FindByID(ctx context.Context, ID string) {
	//
}
