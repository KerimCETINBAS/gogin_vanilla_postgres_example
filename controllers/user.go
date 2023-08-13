package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/services"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type userController struct {
	userService services.IUserService
}
type IUserController interface {
	GetUsers(ctx *gin.Context)
}

func UserController(
	userService services.IUserService) IUserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetUsers(ctx *gin.Context) {

	var (
		err   error
		users *[]types.User
	)

	users, err = c.userService.GetUsers()

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.IndentedJSON(http.StatusOK, users)

}
