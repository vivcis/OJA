package models

import "time"

type User struct {
	ID           string    `json:"id" gorm:"primarykey"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email" gorm:"unique"`
	Username     string    `json:"username" gorm:"unique"`
	PasswordHash string    `json:"-"`
	Address      string    `json:"address"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"password,omitempty" gorm:"-"`
	Reset        string    `json:"-"`
	Image        string    `json:"image"`
	Status       bool      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
