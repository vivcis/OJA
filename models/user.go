package models

type User struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email" gorm:"unique"`
	Username        string `json:"username" gorm:"unique"`
	Password        string `json:"password,omitempty" gorm:"-"`
	ConfirmPassword string `json:"confirm_password,omitempty" gorm:"-"`
	PasswordHash    string `json:"-"`
	Address         string `json:"address"`
	PhoneNumber     string `json:"phone_number"`
	Image           string `json:"image"`
	IsActive        bool   `json:"status"`
	Token           string `json:"token"`
}
