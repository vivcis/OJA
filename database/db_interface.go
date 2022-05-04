package database

import (
	"fmt"
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// DB provides access to the different db
//go:generate  mockgen -source=./db_interface.go -destination=./mockfile.go DB -package=mock
type DB interface {
	//CreateSeller(user *models.Seller) (*models.Seller, error)
	//CreateBuyer(user *models.Buyer) (*models.Buyer, error)
	//FindSellerByUsername(username string) (*models.Seller, error)
	//FindBuyerByUsername(username string) (*models.Buyer, error)
	//FindSellerByEmail(email string) (*models.Seller, error)
	//FindBuyerByEmail(email string) (*models.Buyer, error)
	//FindSellerByPhone(phone string) (*models.Seller, error)
	//FindBuyerByPhone(phone string) (*models.Buyer, error)
	//FindAllSellersExcept(except string) ([]models.Seller, error)
	//UpdateUser(user *models.User) error
	//TokenInBlacklist(token *string) bool
	SearchDB(c *gin.Context) ([]models.Product, error)
}

// ValidationError defines error that occur due to validation
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}

type DBParams struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func InitDBParams() DBParams {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return DBParams{
		Host:     host,
		User:     user,
		Password: password,
		DbName:   dbName,
		Port:     port,
	}
}
