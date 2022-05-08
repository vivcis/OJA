package database

import (
	"errors"
	"fmt"
	"github.com/decadevs/shoparena/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

//PostgresDb implements the DB interface
type PostgresDb struct {
	DB *gorm.DB
}

// Init sets up the mongodb instance
func (pdb *PostgresDb) Init(host, user, password, dbName, port string) error {
	fmt.Println("connecting to Database.....")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if db == nil {
		return fmt.Errorf("database was not initialized")
	} else {
		fmt.Println("Connected to Database")
	}

	err = db.AutoMigrate(&models.Product{}, &models.Seller{}, &models.Buyer{}, &models.Shop{}, &models.Blacklist{})
	if err != nil {
		return fmt.Errorf("migration error: %v", err)
	}

	pdb.DB = db

	return nil

}

// SearchProduct Searches all products from DB
func (pdb *PostgresDb) SearchProduct(lowerPrice, upperPrice, category, name string) ([]models.Product, error) {

	var products []models.Product
	LPInt, _ := strconv.Atoi(lowerPrice)
	UPInt, _ := strconv.Atoi(upperPrice)

	if LPInt == 0 && UPInt == 0 && name == "" {
		err := pdb.DB.Where("product_category = ?", category).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt == 0 && name == "" {
		err := pdb.DB.Where("product_category = ?", category).
			Where("product_price <= ?", uint(UPInt)).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if UPInt == 0 && name == "" {
		err := pdb.DB.Where("product_category = ?", category).
			Where("product_price >= ?", uint(LPInt)).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt != 0 && UPInt != 0 && name == "" {
		err := pdb.DB.Where("product_category = ?", category).Where("product_price >= ?", uint(LPInt)).
			Where("product_price <= ?", uint(UPInt)).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt == 0 && UPInt == 0 && name != "" {
		err := pdb.DB.Where("product_category = ?", category).
			Where("product_name LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt == 0 && name != "" {
		err := pdb.DB.Where("product_category = ?", category).
			Where("product_price <= ?", uint(UPInt)).
			Where("product_name LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if UPInt == 0 && name != "" {
		err := pdb.DB.Where("product_category = ?", category).
			Where("product_price >= ?", uint(LPInt)).
			Where("product_name LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if category == "" {
		return nil, errors.New("choose a category")

	} else {
		err := pdb.DB.Where("product_category = ?", category).Where("product_price >= ?", uint(LPInt)).
			Where("product_price <= ?", uint(UPInt)).
			Where("product_name LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return products, nil
}

// CreateSeller creates a new Seller in the DB
func (pdb *PostgresDb) CreateSeller(user *models.Seller) (*models.Seller, error) {
	_, err := pdb.FindSellerByEmail(user.Email)
	if err == nil {
		return user, ValidationError{Field: "email", Message: "already in use"}
	}
	_, err = pdb.FindSellerByUsername(user.Username)
	if err == nil {
		return user, ValidationError{Field: "username", Message: "already in use"}
	}
	_, err = pdb.FindSellerByPhone(user.PhoneNumber)
	if err == nil {
		return user, ValidationError{Field: "phone", Message: "already in use"}
	}
	user.CreatedAt = time.Now()
	err = pdb.DB.Create(user).Error
	return user, err
}

// CreateBuyer creates a new Buyer in the DB
func (pdb *PostgresDb) CreateBuyer(user *models.Buyer) (*models.Buyer, error) {
	_, err := pdb.FindBuyerByEmail(user.Email)
	if err == nil {
		return user, ValidationError{Field: "email", Message: "already in use"}
	}
	_, err = pdb.FindBuyerByUsername(user.Username)
	if err == nil {
		return user, ValidationError{Field: "username", Message: "already in use"}
	}
	_, err = pdb.FindBuyerByPhone(user.PhoneNumber)
	if err == nil {
		return user, ValidationError{Field: "phone", Message: "already in use"}
	}
	user.CreatedAt = time.Now()
	err = pdb.DB.Create(user).Error
	return user, err
}

// FindSellerByUsername finds a user by the username
func (pdb *PostgresDb) FindSellerByUsername(username string) (*models.Seller, error) {
	user := &models.Seller{}

	if err := pdb.DB.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	if !user.Status {
		return nil, errors.New("user inactive")
	}
	return user, nil
}

// FindBuyerByUsername finds a user by the username
func (pdb *PostgresDb) FindBuyerByUsername(username string) (*models.Buyer, error) {
	buyer := &models.Buyer{}

	if err := pdb.DB.Where("username = ?", username).First(buyer).Error; err != nil {
		return nil, err
	}
	if !buyer.Status {
		return nil, errors.New("user inactive")
	}
	return buyer, nil
}

// FindSellerByEmail finds a user by email
func (pdb *PostgresDb) FindSellerByEmail(email string) (*models.Seller, error) {
	user := &models.Seller{}
	if err := pdb.DB.Where("email = ?", email).First(user).Error; err != nil {
		return nil, errors.New(email + " does not exist" + " user not found")
	}

	return user, nil
}

// FindBuyerByEmail finds a user by email
func (pdb *PostgresDb) FindBuyerByEmail(email string) (*models.Buyer, error) {
	buyer := &models.Buyer{}
	if err := pdb.DB.Where("email = ?", email).First(buyer).Error; err != nil {
		return nil, errors.New(email + " does not exist" + " user not found")
	}

	return buyer, nil
}

// FindSellerByPhone finds a user by the phone
func (pdb PostgresDb) FindSellerByPhone(phone string) (*models.Seller, error) {
	user := &models.Seller{}
	if err := pdb.DB.Where("phone_number =?", phone).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindBuyerByPhone finds a user by the phone
func (pdb PostgresDb) FindBuyerByPhone(phone string) (*models.Buyer, error) {
	buyer := &models.Buyer{}
	if err := pdb.DB.Where("phone_number =?", phone).First(buyer).Error; err != nil {
		return nil, err
	}
	return buyer, nil
}

// TokenInBlacklist checks if token is already in the blacklist collection
func (pdb *PostgresDb) TokenInBlacklist(token *string) bool {
	return false
}

// FindAllUsersExcept returns all the users expcept the one specified in the except parameter
func (pdb *PostgresDb) FindAllSellersExcept(except string) ([]models.Seller, error) {
	sellers := []models.Seller{}
	if err := pdb.DB.Not("username = ?", except).Find(sellers).Error; err != nil {

		return nil, err
	}
	return sellers, nil
}

func (pdb *PostgresDb) UpdateUser(user *models.User) error {

	return nil
}
