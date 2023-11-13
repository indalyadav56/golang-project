package middlewares

import (
	"golang-project/helpers"
	"golang-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func containsPath(urls []string, path string) bool {
	for _, url := range urls {
	  if strings.Contains(url, path) {
		return true
	  }
	}
	return false
  }
  

func AuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		publicUrl := []string{"/v1/auth/login","/v1/users"} // extempted url goes here
		exempted := containsPath(publicUrl,c.Request.URL.Path)
		
		if exempted {
			c.Next()
			return
		}
		
		authHeader := c.Request.Header.Get("Authorization")
		
		if len(strings.TrimSpace(authHeader) ) <= 0 {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Token is required",
				Error: "",
				Data: []interface{}{},
			})
			return
		}

		bearerToken := strings.Split(authHeader, " ")[1]

        err := helpers.VerifyToken(bearerToken)
        
		if err != nil {
           c.JSON(400, map[string]interface{}{
            "status": "Invalid token",
           })
        }
		
		c.Next()
	  }
  
  }