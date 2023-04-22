package helper

import (
	"errors"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	AppJSON                  = "application/json"
	ErrInvalidPhotoId        = "Invalid Photo ID"
	ErrInvalidCommentId      = "Invalid Comment ID"
	ErrInvalidSocialMediaId  = "Invalid Social Media ID"
	ErrNotFound              = "not found"
	DeletePhotoSuccess       = "Photo deleted successfully"
	DeleteCommentSuccess     = "Comment deleted successfully"
	DeleteSocialMediaSuccess = "Social Media deleted successfully"
	MissingUserData          = "missing user data in context"
	InvalidUserData          = "invalid user data type"
	MissingOrInvalidUserId   = "missing or invalid user ID"
)

func GetUserIDFromJwt(c *gin.Context) (uint64, error) {
	userData, ok := c.Get("userData")
	if !ok {
		return 0, errors.New(MissingUserData)
	}
	claims, ok := userData.(jwt.MapClaims)
	if !ok {
		return 0, errors.New(InvalidUserData)
	}
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New(MissingOrInvalidUserId)
	}
	return uint64(id), nil
}

func GetPhotoIdFromContext(c *gin.Context) (uint64, error) {
	return strconv.ParseUint(c.Param("photoId"), 10, 64)
}

func GetCommentIdFromContext(c *gin.Context) (uint64, error) {
	return strconv.ParseUint(c.Param("commentId"), 10, 64)
}

func GetSocialMediaIdFromContext(c *gin.Context) (uint64, error) {
	return strconv.ParseUint(c.Param("socialMediaId"), 10, 64)
}

func BindModelFromContext(c *gin.Context, model interface{}) error {
	contentType := GetContentType(c)

	if contentType == AppJSON {
		if err := c.ShouldBindJSON(model); err != nil {
			return err
		}
	} else {
		if err := c.ShouldBind(model); err != nil {
			return err
		}
	}

	return nil
}

func HandleError(err error, c *gin.Context) {
	if strings.Contains(err.Error(), ErrNotFound) {
		NotFound(c, err.Error())
	} else if strings.Contains(err.Error(), ErrInvalidPhotoId) ||
		strings.Contains(err.Error(), ErrInvalidCommentId) ||
		strings.Contains(err.Error(), ErrInvalidSocialMediaId) {
		BadRequest(c, err.Error())
	} else {
		InternalServerError(c, err.Error())
	}
}

func GetPaginationParams(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit > 1000 {
		limit = 1000
	}
	return page, limit
}
