package http

import (
	"super_admin/handlers"

	"github.com/gofiber/fiber/v2"
)


func UserRouter(app *fiber.App){
	userRouter := app.Group("/users")
	userRouter.Get("", handlers.GetAllUsers)
	userRouter.Get("", handlers.DeleteUser)
}