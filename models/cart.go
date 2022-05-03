package models

type Cart struct {
	ID            uint `gorm:"primarykey"`
	BuyersID      string
	Buyer         Buyer
	Product       []Product `json:"product"`
	TotalPrice    string    `json:"total_price"`
	TotalQuantity string    `json:"total_quantity"`
}
