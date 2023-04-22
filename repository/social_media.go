package repository

import (
	"math"
	"project-final/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// interface SocialMediaRepo
type SocialMediaRepo interface {
	CreateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error)
	GetSocialMediaById(id uint64) (model.SocialMedia, error)
	GetSocialMedias(page int, limit int) (gin.H, error)
	GetSocialMedia(user_id uint64) (model.SocialMedia, error)
	UpdateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error)
	DeleteSocialMedia(user_id uint64) error
}

func (r Repo) CreateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Where(model.SocialMedia{UserID: social_media.UserID}).FirstOrCreate(&social_media).Error
	return social_media, err
}

func (r Repo) GetSocialMediaById(id uint64) (model.SocialMedia, error) {
	var social_media model.SocialMedia
	err := r.db.Preload("User").First(&social_media, id).Error
	return social_media, err
}

func (r Repo) GetSocialMedias(page int, limit int) (gin.H, error) {
	var social_medias []model.SocialMedia
	var count int64
	query := r.db.Preload("User")

	if page == -1 && limit == -1 {
		return r.getAllSocialMedias(query)
	}

	offset := (page - 1) * limit
	r.db.Model(social_medias).Count(&count)

	err := query.Offset(offset).Limit(limit).Find(&social_medias).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"social_medias": social_medias,
		"meta": gin.H{
			"total_data":   count,
			"current_page": page,
			"limit":        limit,
			"total_page":   int(math.Ceil(float64(count) / float64(limit))),
		},
	}

	return responseData, nil
}

func (r Repo) getAllSocialMedias(query *gorm.DB) (gin.H, error) {
	var social_medias []model.SocialMedia

	err := query.Find(&social_medias).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"social_medias": social_medias,
		"meta": gin.H{
			"total_data": len(social_medias),
		},
	}

	return responseData, nil
}

func (r Repo) GetSocialMedia(user_id uint64) (model.SocialMedia, error) {
	var social_media model.SocialMedia
	err := r.db.Where("user_id = ?", user_id).First(&social_media).Error
	return social_media, err
}

func (r Repo) UpdateSocialMedia(social_media model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Updates(&social_media).Error
	return social_media, err
}

func (r Repo) DeleteSocialMedia(user_id uint64) error {
	return r.db.Delete(&model.SocialMedia{}, "user_id = ?", user_id).Error
}
