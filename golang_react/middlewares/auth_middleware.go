package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)


func AuthMiddleware() fiber.Handler  {
   return func(c *fiber.Ctx) error {
        fmt.Println("Received request: ", c.Method(), c.Path())
        err := c.Next()
        fmt.Println("Processed request")
        return err
    }
}
