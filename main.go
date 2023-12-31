package main

import (
	"fmt"

	"golang-project/config"
	"golang-project/middlewares"
	"golang-project/routes"
	"golang-project/websocket"

	_ "golang-project/docs"

	"github.com/gin-contrib/cors"
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

	config.ConnectDB()
    
	server := gin.Default()

	// middlewares
	// Use CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	server.Use(cors.New(config))
	server.Use(middlewares.AuthMiddleware())
	
	// routes
	routes.UserRouter(server)
	routes.AuthRouter(server)
	
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// websocket
	server.GET("/chat", websocket.Websocket)

	server.Run()
	
}

