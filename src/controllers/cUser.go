package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/herEmanuel/go-api/services"
)

//Register controller
func Register(ctx *fiber.Ctx) error {

	var bodyFields map[string]string

	err := ctx.BodyParser(&bodyFields)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).SendString("Could not parse the request body")
	}

	newUser, err := services.Register(bodyFields)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(400).SendString(err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    newUser["token"].(string),
		HTTPOnly: true,
		MaxAge:   60 * 60 * 24 * 7 * 1000,
	})
	return ctx.JSON(newUser["newUser"])
}

//Login controller
func Login(ctx *fiber.Ctx) error {

	var bodyFields map[string]string

	err := ctx.BodyParser(&bodyFields)
	if err != nil {
		return ctx.Status(500).SendString("Could not parse the request body")
	}

	user, err := services.Login(bodyFields)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    user["token"].(string),
		HTTPOnly: true,
		MaxAge:   60 * 60 * 24 * 7 * 1000,
	})
	return ctx.JSON(user["user"])
}

//ChangePassword changes user's password
func ChangePassword(ctx *fiber.Ctx) error {

	var bodyFields map[string]string

	err := ctx.BodyParser(&bodyFields)
	if err != nil {
		return ctx.Status(500).SendString("Could not parse the request body")
	}

	bodyFields["id"] = ctx.Locals("userId").(string)

	err = services.ChangePassword(bodyFields)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.SendString("Your password was changed successfully")
}
