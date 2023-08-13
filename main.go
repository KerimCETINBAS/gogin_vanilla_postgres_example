package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kerimcetinbas/goginpostgrestut/database"
	"github.com/kerimcetinbas/goginpostgrestut/routers"
)

func main() {

	// init database connection
	database.InitDatabase()
	// create router
	router := gin.Default()
	api := router.Group("/api/v1")

	// serve static files
	router.Static("public", "./public")

	// site group
	router.LoadHTMLGlob("template/*/*.html")
	routers.SiteRouter(router)

	// api group
	routers.MessageRouter(api.Group("messages"))
	routers.UserRouter(api.Group("users"))
	routers.AuthRouter(api.Group("auth"))

	// listen on port 8081
	router.Run(":8081")
}
