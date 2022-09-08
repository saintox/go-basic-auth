package services

import "github.com/saintox/go-basic-auth/repositories"

type Service struct {
	Login    LoginService
	Register RegisterService
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Login:    NewLoginService(repository),
		Register: NewRegisterService(repository),
	}
}
