package routes

import (
	"github.com/mgabrielo/golang-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("users/signup", controllers.SignUp())
	route.POST("users/login", controllers.LogIn())
}
