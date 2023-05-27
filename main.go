package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:Sajal@123@/adminapp"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println(db)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Sajal 👋!")
	})

	app.Listen(":8000")
}
