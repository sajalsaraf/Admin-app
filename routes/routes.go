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
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.AllUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
	app.Post("/api/logout", controllers.Logout)
}
