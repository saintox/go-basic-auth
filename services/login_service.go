package services

import (
	"context"
	"github.com/saintox/go-basic-auth/repositories"
)

type LoginService interface {
	FindUserByID(ctx context.Context, ID string)
}

type LoginServiceImpl struct {
	UserRepo repositories.UserRepository
}

func NewLoginService() *LoginServiceImpl {
	return &LoginServiceImpl{
		//UserRepo:
	}
}

func (s LoginServiceImpl) FindUserByID(ctx context.Context, ID string) {
	//
}
