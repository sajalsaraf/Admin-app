package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/database"
	"github.com/sajalsaraf/Admin-app.git/middlewares"
	"github.com/sajalsaraf/Admin-app.git/models"
)

func AllUsers(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page))

}

func CreateUser(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		return err
	}
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("sajal")
	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {

	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id")) // id entered at api will store in this id varible

	user := models.User{
		Id: uint(id), // making a new struct varible and store id in it
	}
	database.DB.Preload("Role").Find(&user) // find func is used to find the entry corresponding to this id and preload func helps to get details of the entered foregin key
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}
func DeleteUser(c *fiber.Ctx) error {
	err := middlewares.IsAuthorized(c, "users")
	if err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	database.DB.Delete(&user)
	return nil
}
