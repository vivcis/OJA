package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	SellerId  uint `json:"seller_id"`
	Seller    Seller
	BuyerId   uint `json:"buyerid"`
	Buyer     Buyer
	ProductId uint `json:"product_id"`
	Product   Product
}
