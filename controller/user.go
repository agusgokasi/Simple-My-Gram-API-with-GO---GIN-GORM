package controller

import (
	"project-final/helper"
	"project-final/model"

	"github.com/gin-gonic/gin"
)

// UserRegister godoc
// @Summary     Register a new user
// @Description Use this api to register a new user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user_request body model.UserRequest true "User request object"
// @Success     200 {object} helper.Response
// @Router      /users/register [post]
func (h HttpServer) UserRegister(c *gin.Context) {
	user := model.User{}
	err := helper.BindModelFromContext(c, &user)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	// call service
	res, err := h.app.CreateUser(user)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// UserLogin godoc
// @Summary     login user
// @Description Use this api to login a user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user_login_request body model.UserLoginRequest true "User request object"
// @Success     200 {object} helper.Response
// @Router      /users/login [post]
func (h HttpServer) UserLogin(c *gin.Context) {
	user := model.User{}
	err := helper.BindModelFromContext(c, &user)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	token, err := h.app.LoginUser(user)
	if err != nil {
		helper.Unauthorized(c, err.Error())
		return
	}
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	helper.Ok(c, token)
}
