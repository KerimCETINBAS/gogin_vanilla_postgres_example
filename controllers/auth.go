package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/services"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type authController struct {
	authSerivce services.IAuthService
	userService services.IUserService
}
type IAuthController interface {
	Signup(ctx *gin.Context)
	Signin(ctx *gin.Context)
}

func AuthController(
	authService services.IAuthService,
	userService services.IUserService,
) IAuthController {
	return &authController{
		authSerivce: authService,
		userService: userService,
	}
}

func (c *authController) Signup(ctx *gin.Context) {
	userCreateDto := types.UserCreateDto{}
	if err := ctx.BindJSON(&userCreateDto); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := c.userService.CreateUser(&userCreateDto); err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
	}
}

func (c *authController) Signin(ctx *gin.Context) {

	var loginDto = types.UserLoginDto{}
	var (
		user types.User
		err  error
	)
	ctx.BindJSON(&loginDto)

	user, err = c.authSerivce.Login(&loginDto)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.SetCookie("session", fmt.Sprintf("%v", user.ID), 3600, "/", "localhost", false, true)
	ctx.AbortWithStatus(http.StatusOK)
}
