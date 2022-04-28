package models

type Shop struct {
	ID       uint   `gorm:"primarykey"`
	ShopName string `json:"shop_name"`
	SellerID string `json:"seller_id"`
	Seller   Seller
	Product  []Product `gorm:"many2many:shop_products;"`
}
