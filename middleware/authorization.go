package middleware

import (
	"project-final/config"
	"project-final/helper"
	"project-final/model"

	"github.com/gin-gonic/gin"
)

var (
	notAllowed = "you're not allowed to access this"
)

func PhotoAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := helper.GetUserIDFromJwt(c)
		if err != nil {
			helper.BadRequest(c, err.Error())
			c.Abort()
			return
		}

		db := config.GetDB()
		id, err := helper.GetPhotoIdFromContext(c)
		if err != nil {
			helper.BadRequest(c, err.Error())
			c.Abort()
			return
		}

		photo := model.Photo{}
		err = db.Select("user_id").First(&photo, id).Error
		if err != nil {
			helper.NotFound(c, err.Error())
			c.Abort()
			return
		}

		if photo.UserID != userID {
			helper.Unauthorized(c, notAllowed)
			c.Abort()
			return
		}

		c.Next()
	}
}

func CommentAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := helper.GetUserIDFromJwt(c)
		if err != nil {
			helper.BadRequest(c, err.Error())
			c.Abort()
			return
		}

		commentID, err := helper.GetCommentIdFromContext(c)
		if err != nil {
			helper.BadRequest(c, err.Error())
			c.Abort()
			return
		}

		db := config.GetDB()
		comment := model.Comment{}
		err = db.Select("user_id").First(&comment, commentID).Error
		if err != nil {
			helper.NotFound(c, err.Error())
			c.Abort()
			return
		}

		if comment.UserID != userID {
			helper.Unauthorized(c, notAllowed)
			c.Abort()
			return
		}

		c.Next()
	}
}

func SocialMediaAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := helper.GetUserIDFromJwt(c)
		if err != nil {
			helper.BadRequest(c, err.Error())
			c.Abort()
			return
		}

		db := config.GetDB()
		socialMedia := model.SocialMedia{}
		err = db.Where("user_id = ?", userID).First(&socialMedia).Error
		if err != nil {
			helper.NotFound(c, err.Error())
			c.Abort()
			return
		}

		c.Set("social_media", socialMedia)
		c.Next()
	}
}
