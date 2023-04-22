package model

import (
	"project-final/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string    `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Your email is not valid"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Your password must be at least 6 characters long"`
	Age      int       `gorm:"not null" json:"age" form:"age" valid:"required~Your Age is required,range(8|100)~Your age must be greater than or equal to 8"`
	Photos   []Photo   `json:"photos,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments []Comment `json:"comments,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UserRequest struct {
	Username string `example:"Agus Setiawan"`
	Email    string `example:"agusgokasi@gmail.com"`
	Password string `example:"123456"`
	Age      int    `example:"25"`
}

type UserLoginRequest struct {
	Email    string `example:"agusgokasi@gmail.com"`
	Password string `example:"123456"`
}

// hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helper.HashPass(u.Password)

	return
}
