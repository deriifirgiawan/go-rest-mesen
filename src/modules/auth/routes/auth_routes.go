package auth

import (
	"rest-app-pos/src/containers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, app *containers.AppContainer) {
	auth := router.Group("/auth")

	{
		auth.POST("/register", app.AuthController.Register)
		auth.POST("/login", app.AuthController.Login)
	}
}