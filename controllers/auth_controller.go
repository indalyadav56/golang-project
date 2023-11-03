package controllers

import "github.com/gin-gonic/gin"


func AuthController(c *gin.Context) {
	c.JSON(200, map[string]string{
		"username": "indalyadv56",
		"password": "Indal@123",
	})
}