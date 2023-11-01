package controllers

import (
	"fmt"
	models "golang-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// func GetUsers(c *gin.Context) {

// 	users, err := services.GetUsers()
// 	if err != nil {
// 		c.JSON(404,map[string]interface{}{"err": err,})
// 		return
// 	}

// 	c.JSON(200, users)
//   }

// func UserRegister(c *gin.Context){

// }

// c.JSON(http.StatusOK, user)
// var user models.User
// Bind and validate the JSON request
// if err := c.ShouldBindJSON(&user); err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	return
// }
// If the request data is valid, generate a new UUID for the user's ID
// user.ID = uuid.New()
// fmt.Println("Incoming request:-", user.UserName, user.Password)
// fmt.Println("User ID:", user.ID)

  type fieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
  }

  func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {

	  var errors []fieldError
	  for _, e := range err.(validator.ValidationErrors) {
		errorMessage := cleanValidationError(e)
		errors = append(errors, fieldError{
		  Field: e.Field(),
		  Error: errorMessage,
		})
	  }

	  c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
	  return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
  }

  // cleanValidationError is a custom function to extract the error message from the ValidationErrors
func cleanValidationError(e validator.FieldError) string {
	// Extract the error message without the package-specific information
	errorMessage := fmt.Sprintf("%s required", e.Field())
	return errorMessage
}