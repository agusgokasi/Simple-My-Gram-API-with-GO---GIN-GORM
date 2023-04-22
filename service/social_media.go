package service

// usecase
import (
	"project-final/model"

	"github.com/gin-gonic/gin"
)

type SocialMediaService interface {
	CreateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error)
	GetSocialMediaById(id uint64) (model.SocialMedia, error)
	GetSocialMedias(page int, limit int) (gin.H, error)
	GetSocialMedia(user_id uint64) (model.SocialMedia, error)
	UpdateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error)
	DeleteSocialMedia(user_id uint64) error
}

func (s *Service) CreateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error) {
	return s.repo.CreateSocialMedia(social_media)
}

func (s *Service) GetSocialMediaById(id uint64) (model.SocialMedia, error) {
	return s.repo.GetSocialMediaById(id)
}

func (s *Service) GetSocialMedias(page int, limit int) (gin.H, error) {
	return s.repo.GetSocialMedias(page, limit)
}

func (s *Service) GetSocialMedia(user_id uint64) (model.SocialMedia, error) {
	return s.repo.GetSocialMedia(user_id)
}

func (s *Service) UpdateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error) {
	return s.repo.UpdateSocialMedia(social_media)
}

func (s *Service) DeleteSocialMedia(user_id uint64) error {
	return s.repo.DeleteSocialMedia(user_id)
}
