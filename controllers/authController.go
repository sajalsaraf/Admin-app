package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/database"
	"github.com/sajalsaraf/Admin-app.git/models"
	"github.com/sajalsaraf/Admin-app.git/util"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}

func Other(c *fiber.Ctx) error {
	return c.SendString("Hello, other!")
}

func Register(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data) // Get data from frontend/postman
	if err != nil {
		fmt.Println("error while receiving data from postman")
	}

	if data["password"] != data["confirm_passowrd"] {
		c.Status(400) //status code -- bad request
		return c.JSON(fiber.Map{
			"msg": "passwords do not match",
		})
	}

	pwd_string := data["password"]
	// password, err := bcrypt.GenerateFromPassword([]byte(pwd_string), 14)
	// if err != nil {
	// 	return c.JSON(fiber.Map{
	// 		"msg": "unable to hash password",
	// 	})
	// }

	user := models.User{}
	user.Firstname = data["firstname"]
	user.Lastname = data["lastname"]
	user.Email = data["email"]
	user.RoleId = 1 // Creates an admin user as for admin role id is 1
	// user.Password = password
	user.SetPassword(pwd_string)

	db := database.DB
	db.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println("error while receiving data from postman")
	}

	user := models.User{}
	db := database.DB
	db.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404) //not found
		return c.JSON(fiber.Map{
			"msg": "entry not found",
		})
	}

	err = user.ComparePassword(data["password"])
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": "wrong password",
		})
	}

	// Created a token which stores the user_id and expires in 24hrs
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id))) // now claims operation in util package

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Stored the token in this cookie which is stored in frontend/postman
	// The cookie is valid for 24 hrs
	// Due to this cookie login is required only once in 24hrs
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	var user models.User
	database.DB.Where("id= ?", id).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		fmt.Println("error while receiving data from postman")
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie) // extract id from cookie

	userId, _ := strconv.Atoi(id) //convert into int

	user := models.User{
		Id:        uint(userId),
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
	}
	database.DB.Model(&user).Updates(data) // update into database

	return c.JSON(user)
}
func UpdatePassword(c *fiber.Ctx) error {
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

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie) // cookie store id in string
	userId, _ := strconv.Atoi(id)  // convert into int
	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["password"])
	database.DB.Model(&user).Updates(data)

	return c.JSON(user)
}
