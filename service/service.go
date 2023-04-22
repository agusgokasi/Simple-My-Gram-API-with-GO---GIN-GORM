package service

import "project-final/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	UserService
	PhotoService
	CommentService
	SocialMediaService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
