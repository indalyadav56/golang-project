package main

import (
	"fmt"

	"golang-project/database"
	"golang-project/routes"
	"golang-project/websocket"

	_ "golang-project/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	database.ConnectDB()
    
	server := gin.Default()
	
	// routes
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.UserRouter(server)

	// websocket
	server.GET("/chat", websocket.Websocket)

	server.Run()
	
}

