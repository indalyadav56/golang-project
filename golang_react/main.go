package main

import (
	"super_admin/config"
	routes "super_admin/http"
	"super_admin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// load environment
	config.LoadEnv()
	
	// initialize database connection
	config.InitDB()

	app := fiber.New()

	app.Use(middlewares.AuthMiddleware())

	routes.UserRouter(app)
	routes.ClientRouter(app)

	// Serve static files from the build folder
	app.Static("/", "./build")

	// Catch-all route for serving index.html
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("build/index.html")
	})

	app.Listen(":8081")
}
