package service

// usecase
import (
	"project-final/model"

	"github.com/gin-gonic/gin"
)

type PhotoService interface {
	CreatePhoto(photo model.Photo) (model.Photo, error)
	GetPhotoById(id uint64) (model.Photo, error)
	GetPhotos(page int, limit int) (gin.H, error)
	UpdatePhoto(photo model.Photo) (model.Photo, error)
	DeletePhoto(id uint64) error
}

func (s *Service) CreatePhoto(photo model.Photo) (model.Photo, error) {
	return s.repo.CreatePhoto(photo)
}

func (s *Service) GetPhotoById(id uint64) (model.Photo, error) {
	return s.repo.GetPhotoById(id)
}

func (s *Service) GetPhotos(page int, limit int) (gin.H, error) {
	return s.repo.GetPhotos(page, limit)
}

func (s *Service) UpdatePhoto(photo model.Photo) (model.Photo, error) {
	return s.repo.UpdatePhoto(photo)
}

func (s *Service) DeletePhoto(id uint64) error {
	return s.repo.DeletePhoto(id)
}
