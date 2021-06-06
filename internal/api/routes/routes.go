package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupRoutes is a route management of the app
func SetupRoutes(r *fiber.App) {
	r.Use(cors.New())
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola team")
	})
	v1 := r.Group("/v1")
	{
		v1.Get("/waves", func(c *fiber.Ctx) error {
			return c.SendString("Hola desde Wavesasas")
		})
	}
}
