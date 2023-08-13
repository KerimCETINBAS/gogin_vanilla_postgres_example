package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/services"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type messageController struct {
	messageService services.IMessageService
}

type IMessageController interface {
	GetMessages(ctx *gin.Context)
	CreateMessage(ctx *gin.Context)
}

func MessageController(messageService services.IMessageService) IMessageController {
	return &messageController{
		messageService: messageService,
	}
}

func (c *messageController) GetMessages(ctx *gin.Context) {

	messages, err := c.messageService.GetMessages()

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

func (c *messageController) CreateMessage(ctx *gin.Context) {
	session, isExist := ctx.Get("session")
	data := types.MessageCreateDto{}
	user := session.(types.User)

	if !isExist {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	err := ctx.BindJSON(&data)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	data.UserId = strconv.FormatUint(uint64(user.ID), 10)

	err = c.messageService.CreateMessage(&data)

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	}

	ctx.Status(http.StatusOK)
}
