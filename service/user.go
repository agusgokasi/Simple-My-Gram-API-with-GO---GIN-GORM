package service

// usecase
import (
	"project-final/model"
)

type UserService interface {
	CreateUser(user model.User) (model.User, error)
	LoginUser(user model.User) (string, error)
}

func (s *Service) CreateUser(user model.User) (model.User, error) {
	return s.repo.CreateUser(user)
}

func (s *Service) LoginUser(user model.User) (string, error) {
	return s.repo.LoginUser(user)
}
