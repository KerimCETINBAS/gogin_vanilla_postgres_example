package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/services"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type siteController struct {
	messageService services.IMessageService
}
type ISiteController interface {
	RenderLoginPage(ctx *gin.Context)
	RenderHomePage(ctx *gin.Context)
	RenderRegisterPage(ctx *gin.Context)
}

func SiteController(messageService services.IMessageService) ISiteController {
	return &siteController{
		messageService: messageService,
	}
}

func (c *siteController) RenderLoginPage(ctx *gin.Context) {
	var (
		messages *[]types.Message
		err      error
	)
	messages, err = c.messageService.GetMessages()

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	fmt.Println(messages)
	ctx.HTML(http.StatusOK, "login", gin.H{
		"title":    "message app",
		"messages": messages,
	})
}

func (c *siteController) RenderHomePage(ctx *gin.Context) {
	session, isSessionExist := ctx.Get("session")

	if !isSessionExist {
		ctx.Redirect(302, "/login")
		return
	}
	var (
		messages *[]types.Message
		err      error
	)
	messages, err = c.messageService.GetMessages()

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.HTML(http.StatusOK, "home", gin.H{
		"title":    "message app",
		"messages": messages,
		"session":  session,
	})
}

func (c *siteController) RenderRegisterPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register", gin.H{
		"title": "message app",
	})
}
