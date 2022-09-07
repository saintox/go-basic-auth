package repositories

import "context"

type UserRepository interface {
	FindByID(ctx context.Context, ID string)
}

type UserRepositoryImpl struct {
	//db *DB
}

func (u UserRepositoryImpl) FindByID(ctx context.Context, ID string) {
	//
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		//
	}
}
