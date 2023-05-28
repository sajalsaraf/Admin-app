package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Get("/first", controllers.Other)
	app.Post("/api/register", controllers.Register)
}
