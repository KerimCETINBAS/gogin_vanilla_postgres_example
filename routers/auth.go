package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/controllers"
	"github.com/kerimcetinbas/goginpostgrestut/services"
)

func AuthRouter(r *gin.RouterGroup) {

	c := controllers.AuthController(
		services.AuthService(),
		services.UserService())
	r.POST("signup", c.Signup)
	r.POST("signin", c.Signin)
}
