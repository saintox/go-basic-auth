package services

import (
	"github.com/saintox/go-basic-auth/entities"
	"github.com/saintox/go-basic-auth/pkg/utils"
	"github.com/saintox/go-basic-auth/repositories"
)

type RegisterService interface {
	UserRegister(data *entities.User) (result *entities.User, err error)
}

type RegisterImpl struct {
	UserRepo repositories.UserRepository
}

func NewRegisterService(repository *repositories.Repository) *RegisterImpl {
	return &RegisterImpl{
		UserRepo: repository.User,
	}
}

func (r RegisterImpl) UserRegister(userData *entities.User) (createdUser *entities.User, err error) {
	hashedPassword, err := utils.HashPassword(userData.Password, 14)
	if err != nil {
		return createdUser, err
	}

	parsedData := &entities.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: hashedPassword,
	}

	createdUser, err = r.UserRepo.CreateUser(parsedData)
	if err != nil {
		return createdUser, err
	}

	return createdUser, err
}
