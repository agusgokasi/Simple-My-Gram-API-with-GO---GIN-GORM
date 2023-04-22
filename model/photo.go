package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoURL string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint64    `json:"user_id"`
	User     *User     `json:"user,omitempty"`
	Comments []Comment `json:"comments,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type PhotoRequest struct {
	Title    string `example:"New Photo Title"`
	Caption  string `example:"New Photo Caption"`
	PhotoURL string `json:"photo_url" example:"https://via.placeholder.com/150/92c952"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return
}
