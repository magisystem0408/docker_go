package main

import (
	"github.com/gofiber/fiber/v2"
	"go-ambassador/src/database"
	"go-ambassador/src/routes"
)

func main() {

	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
