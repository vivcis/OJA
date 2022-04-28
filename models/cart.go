package models

type Cart struct {
	ID            uint `gorm:"primarykey"`
	BuyersID      string
	Buyer         Buyer
	Product       []Product `json:"product"`
	TotalPrice    float64   `json:"total_price"`
	TotalQuantity uint      `json:"total_quantity"`
}
