package services

import (
	"errors"
	"github.com/saintox/go-basic-auth/constants"
	"github.com/saintox/go-basic-auth/dtos"
	"github.com/saintox/go-basic-auth/pkg/utils"
	"github.com/saintox/go-basic-auth/repositories"
)

type LoginService interface {
	CheckCredential(email string, password string) (*dtos.LoginResponse, error)
}

type LoginServiceImpl struct {
	UserRepo repositories.UserRepository
}

func NewLoginService(repository *repositories.Repository) *LoginServiceImpl {
	return &LoginServiceImpl{
		UserRepo: repository.User,
	}
}

func (s LoginServiceImpl) CheckCredential(email string, password string) (data *dtos.LoginResponse, err error) {
	userData, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return data, err
	}

	checkHash := utils.CheckPasswordHash(password, userData.Password)
	if !checkHash {
		return data, errors.New(constants.InvalidCredential)
	}

	data = &dtos.LoginResponse{
		Status: "Logged in",
	}

	return data, nil
}
