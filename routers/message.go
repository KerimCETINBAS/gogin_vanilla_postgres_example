package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/controllers"
	"github.com/kerimcetinbas/goginpostgrestut/middlewares"
	"github.com/kerimcetinbas/goginpostgrestut/repositories"
	"github.com/kerimcetinbas/goginpostgrestut/services"
)

func MessageRouter(r *gin.RouterGroup) {
	fmt.Println("registered")
	c := controllers.MessageController(services.MessageService())

	// get all messages
	r.GET("/", c.GetMessages)

	// create new message if authenticated
	r.POST("/", middlewares.AuthMiddleWare(repositories.UserRepository()), c.CreateMessage)

}
