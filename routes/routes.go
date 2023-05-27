package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/controllers"
)

func SetupGet(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Get("/first", controllers.Other)
}

func SetupPost(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
