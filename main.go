package main

import (
	"fmt"
	"os"

	"golang-project/models"
	routes "golang-project/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the schema to the database
	db.AutoMigrate(&models.User{})
	
	godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
    fmt.Println("Env Data", os.Getenv("DB_USER"))
	server := gin.Default()

	// loging
	server.Use(gin.Logger())

	server.GET("/users/test", func(c *gin.Context) {

		var users []models.User
		db.Find(&users)
		
		c.JSON(200, users)
	  })

	// routes
	routes.UserRouter(server)

	server.Run()
}