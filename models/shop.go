package models

type Shop struct {
	ID       string
	ShopName string `json:"shop_name"`
	SellerID string `json:"seller_id"`
	Seller   Seller
	Product  []Product `gorm:"many2many:shop_products;"`
}
