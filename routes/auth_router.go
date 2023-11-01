package routes

import (
	"golang-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/v1")
	authGroup.GET("/login", controllers.AuthController)
}