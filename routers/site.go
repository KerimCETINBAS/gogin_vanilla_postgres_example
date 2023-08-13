package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/controllers"
	"github.com/kerimcetinbas/goginpostgrestut/middlewares"
	"github.com/kerimcetinbas/goginpostgrestut/repositories"
	"github.com/kerimcetinbas/goginpostgrestut/services"
)

func SiteRouter(r *gin.Engine) {
	c := controllers.SiteController(
		services.MessageService(),
	)

	r.GET("/login", c.RenderLoginPage)
	r.GET("/", middlewares.AuthMiddleWare(repositories.UserRepository()), c.RenderHomePage)
	r.GET("/register", c.RenderRegisterPage)
}
