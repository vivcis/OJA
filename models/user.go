package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"-"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	Password     string `json:"password,omitempty" gorm:"-"`
	Reset        string `json:"-"`
	Image        string `json:"image"`
	Status       bool   `json:"status"`
}

type UpdateUser struct {
	FirstName      string `json:"first_name" binding:"required" form:"first_name"`
	LastName       string `json:"last_name" binding:"required" form:"last_name"`
	PhoneNumber    string `json:"phone" binding:"required" form:"phone1"`
	Email          string `json:"email" binding:"required,email" form:"email"`
	Address        string `json:"address"  form:"address"`
}