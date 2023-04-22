package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint64 `json:"user_id,omitempty"`
	User    *User  `json:"user,omitempty"`
	PhotoID uint64 `json:"photo_id,omitempty"`
	Photo   *Photo `json:"photo,omitempty"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
}

type CommentRequest struct {
	Message string `example:"Test New Comment"`
}

// hooks
func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return
}
