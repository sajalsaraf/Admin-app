package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sajalsaraf/Admin-app.git/database"
	"github.com/sajalsaraf/Admin-app.git/models"
	"golang.org/x/crypto/bcrypt"
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

	pwd_string := data["password"]
	password, err := bcrypt.GenerateFromPassword([]byte(pwd_string), 14)
	if err != nil {
		return c.JSON(fiber.Map{
			"msg": "unable to hash password",
		})
	}

	user := models.User{}
	user.Firstname = data["firstname"]
	user.Lastname = data["lastname"]
	user.Email = data["email"]
	user.Password = password

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
		c.Status(404)
		return c.JSON(fiber.Map{
			"msg": "entry not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": "wrong password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		// Issuer:    user.Firstname+" "+user.Lastname,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(token)

}
