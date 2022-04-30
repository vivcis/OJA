package database

import (
	"fmt"

	"github.com/decadevs/shoparena/models"
)

// DB provides access to the different db
type DB interface {
	CreateSeller(user *models.Seller) (*models.Seller, error)
	CreateBuyer(user *models.Buyer) (*models.Buyer, error)
	FindSellerByUsername(username string) (*models.Seller, error)
	FindBuyerByUsername(username string) (*models.Buyer, error)
	FindSellerByEmail(email string) (*models.Seller, error)
	FindBuyerByEmail(email string) (*models.Buyer, error)
	FindSellerByPhone(phone string) (*models.Seller, error)
	FindBuyerByPhone(phone string) (*models.Buyer, error)
	FindAllSellersExcept(except string) ([]models.Seller, error)
	UpdateUser(user *models.User) error
	TokenInBlacklist(token *string) bool
}

// ValidationError defines error that occur due to validation
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}
