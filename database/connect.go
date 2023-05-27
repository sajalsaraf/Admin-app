package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:Sajal@123@/adminapp"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	fmt.Println("sajal")
}
