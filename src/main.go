package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/herEmanuel/go-api/database"
	"github.com/herEmanuel/go-api/models"
	"github.com/herEmanuel/go-api/router"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not load the environment variables")
	}

	database.Db, err = gorm.Open(postgres.Open("host=localhost dbname="+os.Getenv("DB_NAME")+" port=5432 user=postgres password="+os.Getenv("DB_PASSWORD")), &gorm.Config{})
	if err != nil {
		log.Fatalln("Could not connect to the database")
	}

	database.Db.AutoMigrate(&models.User{}, &models.Product{}, &models.Purchase{}, &models.Commentary{}, &models.Category{}, &models.CartProduct{})

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
