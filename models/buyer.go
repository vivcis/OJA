package models

type Buyer struct {
	User
	Product []Product `gorm:"many2many:buyer_products;"`
}
