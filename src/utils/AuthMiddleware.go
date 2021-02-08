package utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {

	authToken := ctx.Cookies("auth")
	if authToken == "" {
		return ctx.Status(401).SendString("Invalid token")
	}

	userID, err := VerifyToken(authToken, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	ctx.Locals("userId", userID)

	return ctx.Next()
}
