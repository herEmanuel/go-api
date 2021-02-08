package controllers 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetSingleProduct(ctx *fiber.Ctx) error {
	productId := ctx.Params("id")
	fmt.Println(productId)
	return ctx.SendString("Success")
}