package routes

import (
	"github.com/mgabrielo/golang-auth/controllers"

	middleware "github.com/mgabrielo/golang-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate())
	routes.GET("/users", controllers.GetUsers())
	routes.GET("/users/:user_id", controllers.GetUser())
}
