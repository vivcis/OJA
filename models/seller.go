package models

type Seller struct {
	User
	Rating string `json:"rating"`
}
