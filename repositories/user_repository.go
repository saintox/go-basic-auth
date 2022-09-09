package repositories

import (
	"github.com/saintox/go-basic-auth/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (user entities.User, err error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) FindByEmail(email string) (user entities.User, err error) {
	result := u.db.First(&user, "email = ?", email)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return user, result.Error
	}

	return user, nil
}
