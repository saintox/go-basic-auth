package services

import (
	"context"
	"github.com/saintox/go-basic-auth/entities"
	"github.com/saintox/go-basic-auth/repositories"
)

type RegisterService interface {
	UserRegister(ctx context.Context, data entities.User)
}

type RegisterImpl struct {
	UserRepo repositories.UserRepository
}

func NewRegisterService(repository *repositories.Repository) *RegisterImpl {
	return &RegisterImpl{
		UserRepo: repository.User,
	}
}

func (r RegisterImpl) UserRegister(ctx context.Context, data entities.User) {
	//
}
