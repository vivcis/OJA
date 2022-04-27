package models

type Seller struct {
	User
	Rating int `json:"rating"`
}
