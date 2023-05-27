package database

import (
	"fmt"

	"github.com/sajalsaraf/Admin-app.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	database, err := gorm.Open(mysql.Open("root:Sajal@123@/adminapp"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println("sajal")

	database.AutoMigrate(&models.User{})
}
