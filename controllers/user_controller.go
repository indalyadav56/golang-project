package controllers

import (
	"golang-project/models"
	"golang-project/services"
	"golang-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {

	users, err := services.GetUsers()
	if err != nil {
		c.JSON(404,map[string]interface{}{"err": err,})
		return
	}

	c.JSON(200, users)
  }


  func CreateUser(c *gin.Context) {
	  var user models.User
	  if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	hashedPassword, err := hashPassword(user.Password)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword
	userData, err := services.CreateUser(user)

	if err!= nil{
		c.JSON(253, gin.H{"error": err.Error()})
	}

	// Generate and set JWT token
	token, _ := utils.GenerateToken(user)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 3600 * 24 * 7, "", "", false, true)


	c.JSON(http.StatusCreated, gin.H{"data": userData, "token": token})
  }


  func UpdateUser(c *gin.Context) {
	  var user models.User
	  if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, _ := services.UpdateUser(user)
	c.JSON(http.StatusCreated, userData)
  }

  
  func DeleteUser(c *gin.Context) {
	  var user models.User
	  if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userData, _ := services.DeleteUser(user)
	c.JSON(http.StatusCreated, userData)
  }


  func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


//   type fieldError struct {
// 	Field string `json:"field"`
// 	Error string `json:"error"`
//   }
//   func CreateUser2(c *gin.Context) {
// 	var user models.User

// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 	  var errors []fieldError
// 	  for _, e := range err.(validator.ValidationErrors) {
// 		errorMessage := cleanValidationError(e)
// 		errors = append(errors, fieldError{
// 		  Field: e.Field(),
// 		  Error: errorMessage,
// 		})
// 	  }

// 	  c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
// 	  return
// 	}
// 	DB
// 	c.JSON(http.StatusOK, gin.H{"message": "User created"})
//   }


  // cleanValidationError is a custom function to extract the error message from the ValidationErrors
// func cleanValidationError(e validator.FieldError) string {
	// Extract the error message without the package-specific information
	// errorMessage := fmt.Sprintf("%s required", e.Field())
// 	return errorMessage
// }


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
