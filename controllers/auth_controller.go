package controllers

import (
	"golang-project/config"
	"golang-project/models"
	"golang-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthController(c *gin.Context) {
	 var loginResponse LoginResponse
	 
	 if err := c.ShouldBindJSON(&loginResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }

	  // Find user with provided email
	 var user models.User
	 if err := config.DB.Where("email = ?", loginResponse.Email).First(&user).Error; err != nil {
	   c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	   return
	 }

	  // Check if provided password matches user's password
	  if err := models.VerifyPassword(user.Password, loginResponse.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"}) 
		return
	  }

	// Generate and set JWT token
	token, _ := utils.GenerateToken(user)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 3600 * 24 * 7, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"data": user, "token": token})
}