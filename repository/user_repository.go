package repository

import (
	"gorm.io/gorm"
	"miniproject-3/model/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindByUsername(userUsername string) (domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
	err := repository.DB.Create(&user).Error
	return user, err
}

func (repository *UserRepositoryImpl) FindByUsername(userUsername string) (domain.User, error) {
	var user domain.User
	err := repository.DB.Take(&user, "username = ?", userUsername).Error
	return user, err
}
