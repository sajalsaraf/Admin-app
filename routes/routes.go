package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/controllers"
	"github.com/sajalsaraf/Admin-app.git/middlewares"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Get("/first", controllers.Other)
	app.Post("/api/register", controllers.Register)

	app.Use(middlewares.IsAuthenticated)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
