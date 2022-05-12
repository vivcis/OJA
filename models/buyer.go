package models

import "gorm.io/gorm"

type Buyer struct {
	gorm.Model
	User
	Cart   Cart    `json:"cart"`
	Orders []Order `json:"orders" gorm:"oneToMany"`
}
