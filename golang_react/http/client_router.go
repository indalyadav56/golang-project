package http

import (
	"super_admin/handlers"

	"github.com/gofiber/fiber/v2"
)


func ClientRouter(app *fiber.App){
	clientRouter := app.Group("/clients")
	
	clientRouter.Get("", handlers.GetAllClients)
	clientRouter.Delete("/:clientId", handlers.DeleteClient)
}