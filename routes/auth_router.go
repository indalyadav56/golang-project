package routes

import (
	"golang-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	authGroup := r.Group("/v1/auth")
	authGroup.POST("/login", controllers.AuthController)
}