package main

import (
	"github.com/herEmanuel/go-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Teste")
	})
	router.SetupRoutes(app)
	app.Listen(":3000")
}
