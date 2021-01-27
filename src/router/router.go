package router

import (
	"github.com/herEmanuel/go-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/product", controllers.GetProduct)
}
