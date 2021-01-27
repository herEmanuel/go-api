package main

import (
	"github.com/herEmanuel/go-api/src/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)
	app.Listen(":3000")
}
