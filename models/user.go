package models

type User struct {
	Id        uint
	Firstname string
	Lastname  string
	Email     string `gorm:"unique"`
	Password  []byte
}
