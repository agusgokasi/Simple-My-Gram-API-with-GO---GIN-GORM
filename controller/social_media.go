package controller

import (
	"project-final/helper"
	"project-final/model"

	"github.com/gin-gonic/gin"
)

// CreateSocialMedia godoc
// @Summary 	Create a new social media account
// @Description Use this API to create a new social media account (need authorization bearer token in headers)
// @Tags 		social media
// @Accept 		json
// @Produce 	json
// @Param 		social_media_request body model.SocialMediaRequest true "Social media request object"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/social-media [post]
func (h HttpServer) CreateSocialMedia(c *gin.Context) {
	userID, err := helper.GetUserIDFromJwt(c)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	social_media := model.SocialMedia{}
	err = helper.BindModelFromContext(c, &social_media)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	social_media.UserID = userID
	// call service
	res, err := h.app.CreateSocialMedia(social_media)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetSocialMediaByID godoc
// @Summary 	Show a social media account
// @Description Use this API to get detail social media account by ID (need authorization bearer token in headers)
// @Tags 		social media
// @Accept 		json
// @Produce 	json
// @Param 		socialMediaId path int true "Social Media ID"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/social-media/{socialMediaId} [get]
func (h HttpServer) GetSocialMediaByID(c *gin.Context) {
	id, err := helper.GetSocialMediaIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidSocialMediaId)
		return
	}

	// call service
	res, err := h.app.GetSocialMediaById(id)
	if err != nil {
		helper.HandleError(err, c)
		return
	}

	helper.Ok(c, res)
}

// GetAllSocialMedia godoc
// @Summary 	Get all social media accounts
// @Description Use this API to get all social media accounts with pagination (need authorization bearer token in headers)
// @Tags 		social media
// @Accept 		json
// @Produce 	json
// @Param 		page query int false "Page number, set -1 for all pages"
// @Param 		limit query int false "Limit per page, set -1 for unlimited"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth
// @Router 		/social-media [get]
func (h HttpServer) GetAllSocialMedia(c *gin.Context) {
	page, limit := helper.GetPaginationParams(c)
	// call service
	res, err := h.app.GetSocialMedias(page, limit)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// UpdateSocialMedia godoc
// @Summary 	Update an existing social media account
// @Description Use this API to update an existing social media account by ID (need authorization bearer token in headers and social media token in headers)
// @Tags 		social media
// @Accept 		json
// @Produce 	json
// @Param 		social_media_request body model.SocialMediaRequest true "Social media request object"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth, socialMediaTokenAuth
// @Router 		/social-media [put]
func (h HttpServer) UpdateSocialMedia(c *gin.Context) {
	userID, err := helper.GetUserIDFromJwt(c)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	social_media, err := h.app.GetSocialMedia(userID)
	if err != nil {
		helper.HandleError(err, c)
		return
	}

	social_media_id := social_media.ID

	social_media = model.SocialMedia{}
	err = helper.BindModelFromContext(c, &social_media)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	social_media.ID = social_media_id
	social_media.UserID = userID
	// call service
	res, err := h.app.UpdateSocialMedia(social_media)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// DeleteSocialMedia godoc
// @Summary 	Delete an existing social media account
// @Description Use this API to delete an existing social media account by ID (need authorization bearer token in headers and social media token in headers)
// @Tags 		social media
// @Accept 		json
// @Produce 	json
// @Param 		socialMediaId path int true "Social Media ID"
// @Success 	200 {object} helper.Response
// @Security 	bearerAuth, socialMediaTokenAuth
// @Router 		/social-media [delete]
func (h HttpServer) DeleteSocialMedia(c *gin.Context) {
	userID, err := helper.GetUserIDFromJwt(c)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	err = h.app.DeleteSocialMedia(userID)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, helper.DeleteSocialMediaSuccess)
}
