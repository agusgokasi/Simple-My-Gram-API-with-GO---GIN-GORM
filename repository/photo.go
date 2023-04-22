package repository

import (
	"math"
	"project-final/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// interface PhotoRepo
type PhotoRepo interface {
	CreatePhoto(photo model.Photo) (model.Photo, error)
	GetPhotoById(id uint64) (model.Photo, error)
	GetPhotos(page int, limit int) (gin.H, error)
	UpdatePhoto(photo model.Photo) (model.Photo, error)
	DeletePhoto(id uint64) error
}

func (r Repo) CreatePhoto(photo model.Photo) (model.Photo, error) {
	err := r.db.Create(&photo).Error
	return photo, err
}

func (r Repo) GetPhotoById(id uint64) (model.Photo, error) {
	var photo model.Photo
	err := r.db.Preload("User").First(&photo, id).Error
	return photo, err
}

func (r Repo) GetPhotos(page int, limit int) (gin.H, error) {
	var photos []model.Photo
	var count int64
	query := r.db.Preload("User")

	if page == -1 && limit == -1 {
		return r.getAllPhotos(query)
	}

	offset := (page - 1) * limit
	r.db.Model(photos).Count(&count)

	err := query.Offset(offset).Limit(limit).Find(&photos).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"photos": photos,
		"meta": gin.H{
			"total_data":   count,
			"current_page": page,
			"limit":        limit,
			"total_page":   int(math.Ceil(float64(count) / float64(limit))),
		},
	}

	return responseData, nil
}

func (r Repo) getAllPhotos(query *gorm.DB) (gin.H, error) {
	var photos []model.Photo

	err := query.Find(&photos).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"photos": photos,
		"meta": gin.H{
			"total_data": len(photos),
		},
	}

	return responseData, nil
}

func (r Repo) UpdatePhoto(photo model.Photo) (model.Photo, error) {
	err := r.db.Updates(&photo).Error
	return photo, err
}

func (r Repo) DeletePhoto(id uint64) error {
	return r.db.Delete(&model.Photo{}, "id = ?", id).Error
}
