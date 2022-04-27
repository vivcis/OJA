package models

import "time"

type Product struct {
	ID 				string 		`json:"id" gorm:"primarykey"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	ShopName    	string `json:"shop_name"`
	ProductName     string    `json:"product_name"`
	ProductPrice    float32   `json:"product_price"`
	ProductCategory string    `json:"product_category"`
	ProductImage    string    `json:"product_image"`
	ProductDetails  string    `json:"product_details"`
	Rating          int       `json:"rating"`
	Quantity        int       `json:"quantity"`
}