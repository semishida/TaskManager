package models

type User struct {
	Email    string `gorm:"unique"`
	Password string
	Task
}
