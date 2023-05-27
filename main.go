package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/database"
	"github.com/sajalsaraf/Admin-app.git/routes"
)

func main() {

	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
