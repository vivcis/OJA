package models

type Cart struct {
	ID            uint `gorm:"primarykey"`
	BuyersID      uint
	Buyer         Buyer
	Product       []Product `json:"product"`
	TotalPrice    uint      `json:"total_price"`
	TotalQuantity uint      `json:"total_quantity"`
}
