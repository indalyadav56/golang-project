package routes

import (
	userController "golang-project/controllers"

	"github.com/gin-gonic/gin"
)


func UserRouter(r *gin.Engine){
	userGroup:= r.Group("/v1/users")

	userGroup.POST("", userController.CreateUser)
	userGroup.GET("", userController.GetUsers)
	userGroup.PUT("", userController.UpdateUser)
	userGroup.DELETE("", userController.DeleteUser)
	
}