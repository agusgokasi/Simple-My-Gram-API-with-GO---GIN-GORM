package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required"`
	UserID         uint64
	User           *User `json:"User,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type SocialMediaRequest struct {
	Name           string `example:"Agus Gokasi"`
	SocialMediaURL string `json:"social_media_url" example:"https://www.facebook.com/agusgokasi1/"`
}

// hooks
func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}

	return
}
