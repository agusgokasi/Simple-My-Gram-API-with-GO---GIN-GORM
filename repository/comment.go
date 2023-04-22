package repository

import (
	"math"
	"project-final/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// interface CommentRepo
type CommentRepo interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetComments(photo_id uint64, page int, limit int) (gin.H, error)
	GetCommentById(id uint64) (model.Comment, error)
	UpdateComment(comment model.Comment) (model.Comment, error)
	DeleteComment(id uint64) error
}

func (r Repo) CreateComment(comment model.Comment) (model.Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r Repo) GetComments(photo_id uint64, page int, limit int) (gin.H, error) {
	var comment []model.Comment
	var count int64
	query := r.db.Preload("User").Where("photo_id = ?", photo_id)

	if page == -1 && limit == -1 {
		return r.getAllComment(query)
	}

	offset := (page - 1) * limit
	r.db.Model(comment).Count(&count)

	err := query.Offset(offset).Limit(limit).Find(&comment).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"comment": comment,
		"meta": gin.H{
			"total_data":   count,
			"current_page": page,
			"limit":        limit,
			"total_page":   int(math.Ceil(float64(count) / float64(limit))),
		},
	}

	return responseData, nil
}

func (r Repo) getAllComment(query *gorm.DB) (gin.H, error) {
	var comment []model.Comment

	err := query.Find(&comment).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"comment": comment,
		"meta": gin.H{
			"total_data": len(comment),
		},
	}

	return responseData, nil
}

func (r Repo) GetCommentById(id uint64) (model.Comment, error) {
	var comment model.Comment
	err := r.db.Preload("User").First(&comment, id).Error
	return comment, err
}

func (r Repo) UpdateComment(comment model.Comment) (model.Comment, error) {
	err := r.db.Updates(&comment).Error
	return comment, err
}

func (r Repo) DeleteComment(id uint64) error {
	return r.db.Delete(&model.Comment{}, "id = ?", id).Error
}
