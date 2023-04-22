package repository

import (
	"project-final/helper"
	"project-final/model"
)

// clean architectures -> handler->service->repo

// interface User
type UserRepo interface {
	CreateUser(user model.User) (model.User, error)
	LoginUser(user model.User) (string, error)
}

func (r Repo) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r Repo) LoginUser(user model.User) (string, error) {
	passwordClient := user.Password
	result := r.db.Where("email = ?", user.Email).Take(&user)
	if result.Error != nil {
		return "", errInvalidUser
	}

	if !helper.ComparePass([]byte(user.Password), []byte(passwordClient)) {
		return "", errInvalidUser
	}

	token := helper.GenerateToken(user.ID, user.Email)

	return token, nil
}
