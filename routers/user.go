package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/controllers"
	"github.com/kerimcetinbas/goginpostgrestut/services"
)

func UserRouter(r *gin.RouterGroup) {
	c := controllers.UserController(services.UserService())

	r.GET("/", c.GetUsers)
}
