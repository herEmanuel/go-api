package controllers

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/herEmanuel/go-api/services"
)

func CreateProduct(ctx *fiber.Ctx) error {

	var bodyFields map[string]interface{}

	err := ctx.BodyParser(&bodyFields)
	if err != nil {
		return ctx.Status(500).SendString("Could not parse the request body")
	}

	product, err := services.CreateProduct(bodyFields)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.JSON(product)
}

func AddImage(ctx *fiber.Ctx) error {

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(500).SendString("Could not parse the multipart form")
	}

	images := form.File["productImage"]

	for _, image := range images {

		if image.Size > 1024*1024*3 {
			return ctx.Status(400).SendString("The image can not be bigger than 3 mb")
		}

		allowed := false
		for _, imgType := range []string{"image/png", "image/jpeg"} {
			if imgType == image.Header["Content-Type"][0] {
				allowed = true
			}
		}

		if !allowed {
			return ctx.Status(400).SendString("Invalid image type")
		}

		imageName := fmt.Sprintf("%s$%v%s", strings.TrimSuffix(image.Filename, filepath.Ext(image.Filename)), time.Now().Unix(), filepath.Ext((image.Filename)))

		err = ctx.SaveFile(image, fmt.Sprintf("../uploads/%s", imageName))
		if err != nil {
			fmt.Println(err.Error())
			return ctx.Status(500).SendString("Could not upload the image")
		}

	}

	return ctx.SendString("Image added successfully")
}
