package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ProductId uint   `json:"product_id"`
	Url       string `json:"url"`
}
