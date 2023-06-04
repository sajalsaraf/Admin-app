package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/database"
	"github.com/sajalsaraf/Admin-app.git/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
