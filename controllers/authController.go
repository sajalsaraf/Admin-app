package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/models"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}

func Other(c *fiber.Ctx) error {
	return c.SendString("Hello, other!")
}

func Register(c *fiber.Ctx) error {
	user := models.User{}
	user.Firstname = "Sajal"
	user.Lastname = "Saraf"
	user.Email = "sajalsaraf161203@gmail.com"
	return c.JSON(user)
}
