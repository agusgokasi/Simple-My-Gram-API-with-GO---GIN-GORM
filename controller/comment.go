package controller

import (
	"fmt"
	"project-final/helper"
	"project-final/model"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) GetPhotoFromContext(c *gin.Context) (*model.Photo, error) {
	id, err := helper.GetPhotoIdFromContext(c)
	if err != nil {
		return nil, fmt.Errorf(helper.ErrInvalidPhotoId)
	}

	photo, err := h.app.GetPhotoById(id)
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

// CreateComment godoc
// @Summary		Create a new comment on a photo
// @Description Use this API to create a new comment on a photo (need authorization bearer token in headers)
// @Tags		comments
// @Accept		json
// @Produce 	json
// @Param 		photoId path int true "Photo ID"
// @Param 		comment_request body model.CommentRequest true "Comment request object"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/photos/{photoId}/comments [post]
func (h HttpServer) CreateComment(c *gin.Context) {
	photo, err := h.GetPhotoFromContext(c)
	if err != nil {
		helper.HandleError(err, c)
		return
	}

	comment := model.Comment{}
	err = helper.BindModelFromContext(c, &comment)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	comment.PhotoID = photo.ID
	comment.UserID = photo.UserID
	// call service
	res, err := h.app.CreateComment(comment)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetAllComment godoc
// @Summary 	Get all comments on a photo
// @Description Use this API to get all comments on a photo with pagination (need authorization bearer token in headers)
// @Tags 		comments
// @Accept 		json
// @Produce 	json
// @Param 		photoId path int true "Photo ID"
// @Param 		page query int false "Page number, set -1 for all pages"
// @Param 		limit query int false "Limit per page, set -1 for unlimited"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/photos/{photoId}/comments [get]
func (h HttpServer) GetAllComment(c *gin.Context) {
	id, err := helper.GetPhotoIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidPhotoId)
		return
	}

	page, limit := helper.GetPaginationParams(c)
	// call service
	res, err := h.app.GetComments(id, page, limit)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetCommentByID godoc
// @Summary 	Show a comment by ID
// @Description Use this API to get detail comment by ID (need authorization bearer token in headers)
// @Tags 		comments
// @Accept 		json
// @Produce 	json
// @Param 		commentId path int true "Comment ID"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/comments/{commentId} [get]
func (h HttpServer) GetCommentByID(c *gin.Context) {
	id, err := helper.GetCommentIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidCommentId)
		return
	}

	// call service
	res, err := h.app.GetCommentById(id)
	if err != nil {
		helper.HandleError(err, c)
		return
	}

	helper.Ok(c, res)
}

// UpdateComment godoc
// @Summary 	Update an existing comment by ID
// @Description Use this API to update an existing comment by ID (need authorization bearer token in headers and comment token in headers)
// @Tags 		comments
// @Accept 		json
// @Produce 	json
// @Param 		commentId path int true "Comment ID"
// @Param 		comment_request body model.CommentRequest true "Comment request object"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth, commentTokenAuth
// @Router 		/comments/{commentId} [put]
func (h HttpServer) UpdateComment(c *gin.Context) {
	id, err := helper.GetCommentIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidCommentId)
		return
	}

	comment := model.Comment{}
	err = helper.BindModelFromContext(c, &comment)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	userID, err := helper.GetUserIDFromJwt(c)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	comment.ID = id
	comment.UserID = userID

	// call service
	res, err := h.app.UpdateComment(comment)
	if err != nil {
		helper.HandleError(err, c)
		return
	}

	helper.Ok(c, res)
}

// DeleteComment godoc
// @Summary 	Delete an existing comment by ID
// @Description Use this API to delete an existing comment by ID (need authorization bearer token in headers and comment token in headers)
// @Tags 		comments
// @Accept 		json
// @Produce 	json
// @Param 		commentId path int true "Comment ID"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth, commentTokenAuth
// @Router 		/comments/{commentId} [delete]
func (h HttpServer) DeleteComment(c *gin.Context) {
	id, err := helper.GetCommentIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidPhotoId)
		return
	}

	// call service
	err = h.app.DeleteComment(id)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, helper.DeleteCommentSuccess)
}
