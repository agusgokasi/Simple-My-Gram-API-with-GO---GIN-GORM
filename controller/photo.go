package controller

import (
	"project-final/helper"
	"project-final/model"

	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
// @Summary     Create a new photo
// @Description Use this API to create a new photo (need authorization bearer token in headers)
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       photo_request body model.PhotoRequest true "Photo request object"
// @Success     200 {object} helper.Response
// @Security    bearerAuth
// @Router      /photos [post]
func (h HttpServer) CreatePhoto(c *gin.Context) {
	photo := model.Photo{}
	err := helper.BindModelFromContext(c, &photo)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	userID, err := helper.GetUserIDFromJwt(c)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	photo.UserID = userID
	// call service
	res, err := h.app.CreatePhoto(photo)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetPhotoByID godoc
// @Summary     Show a photo
// @Description Use this API to get detail photo by id (need authorization bearer token in headers)
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id path int true "Photo ID"
// @Success     200 {object} helper.Response
// @Security    bearerAuth
// @Router      /photos/{id} [get]
func (h HttpServer) GetPhotoByID(c *gin.Context) {
	id, err := helper.GetPhotoIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidPhotoId)
		return
	}

	// call service
	res, err := h.app.GetPhotoById(id)
	if err != nil {
		helper.HandleError(err, c)
		return
	}

	helper.Ok(c, res)
}

// GetAllPhoto godoc
// @Summary		Get all photos
// @Description Use this API to get all photos with pagination (need authorization bearer token in headers)
// @Tags 		photos
// @Accept 		json
// @Produce 	json
// @Param 		page query int false "Page number, set -1 for all pages"
// @Param 		limit query int false "Limit per page, set -1 for unlimited"
// @Success		200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/photos [get]
func (h HttpServer) GetAllPhoto(c *gin.Context) {
	page, limit := helper.GetPaginationParams(c)
	// call service
	res, err := h.app.GetPhotos(page, limit)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// UpdatePhoto godoc
// @Summary 	Update an existing photo
// @Description Use this API to update an existing photo by ID (need authorization bearer token in headers)
// @Tags 		photos
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "Photo ID"
// @Param 		photo_request body model.PhotoRequest true "Photo request object"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/photos/{id} [put]
func (h HttpServer) UpdatePhoto(c *gin.Context) {
	id, err := helper.GetPhotoIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidPhotoId)
		return
	}

	photo := model.Photo{}
	err = helper.BindModelFromContext(c, &photo)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	userID, err := helper.GetUserIDFromJwt(c)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	photo.ID = id
	photo.UserID = userID
	// call service
	res, err := h.app.UpdatePhoto(photo)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// DeletePhoto godoc
// @Summary 	Delete an existing photo
// @Description Use this API to delete an existing photo by ID (need authorization bearer token in headers)
// @Tags 		photos
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "Photo ID"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/photos/{id} [delete]
func (h HttpServer) DeletePhoto(c *gin.Context) {
	id, err := helper.GetPhotoIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidPhotoId)
		return
	}

	// call service
	err = h.app.DeletePhoto(id)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, helper.DeletePhotoSuccess)
}
