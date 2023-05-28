package controllers

import (
	"fmt"

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

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println("error while receiving data from postman")
	}

	if data["password"] != data["confirm_passowrd"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": "passwords do not match",
		})
	}

	user := models.User{}
	user.Firstname = data["first_name"]
	user.Lastname = data["last_name"]
	user.Email = data["email_id"]
	user.Password = data["password"]

	return c.JSON(user)
}
