package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herEmanuel/go-api/controllers"
	"github.com/herEmanuel/go-api/utils"
)

//SetupRoutes sets up all the routes in the application
func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	v1.Get("/product/:id", utils.AuthMiddleware, controllers.GetSingleProduct)
	v1.Post("/register", controllers.Register)
	v1.Post("/login", controllers.Login)
	v1.Post("/addImage", controllers.AddImage)
}
