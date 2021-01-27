package controllers

import "github.com/gofiber/fiber/v2"

func GetProduct(context *fiber.Ctx) error {
	return context.SendString("Got the product")
}
