package service

// usecase
import (
	"project-final/model"

	"github.com/gin-gonic/gin"
)

type CommentService interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetComments(photo_id uint64, page int, limit int) (gin.H, error)
	GetCommentById(id uint64) (model.Comment, error)
	UpdateComment(comment model.Comment) (model.Comment, error)
	DeleteComment(id uint64) error
}

func (s *Service) CreateComment(comment model.Comment) (model.Comment, error) {
	return s.repo.CreateComment(comment)
}

func (s *Service) GetComments(photo_id uint64, page int, limit int) (gin.H, error) {
	return s.repo.GetComments(photo_id, page, limit)
}

func (s *Service) GetCommentById(id uint64) (model.Comment, error) {
	return s.repo.GetCommentById(id)
}

func (s *Service) UpdateComment(comment model.Comment) (model.Comment, error) {
	return s.repo.UpdateComment(comment)
}

func (s *Service) DeleteComment(id uint64) error {
	return s.repo.DeleteComment(id)
}
