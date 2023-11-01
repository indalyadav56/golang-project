package middlewares

import (
	"strings"

	helper "golang-project/helpers"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		bearerToken := strings.Split(authHeader, " ")[1]
        err := helper.VerifyToken(bearerToken)
        if err != nil {
           c.JSON(400, map[string]interface{}{
            "status": "Invalid token",
           })
        }
		c.Next()
	  }
  
  }