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

	err = db.AutoMigrate(&models.Category{}, &models.Seller{}, &models.Product{}, &models.Image{},
		&models.Buyer{}, &models.Cart{}, &models.CartProduct{}, &models.Order{}, &models.Blacklist{})
	if err != nil {
		return fmt.Errorf("migration error: %v", err)
	}

	pdb.DB = db

	return nil

}

// SearchProduct Searches all products from DB
func (pdb *PostgresDb) SearchProduct(lowerPrice, upperPrice, categoryName, name string) ([]models.Product, error) {
	categories := models.Category{}
	var products []models.Product

	LPInt, _ := strconv.Atoi(lowerPrice)
	UPInt, _ := strconv.Atoi(upperPrice)

	if categoryName == "" {
		err := pdb.DB.Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return products, nil
	} else {
		err := pdb.DB.Where("name = ?", categoryName).First(&categories).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	category := categories.ID

	if LPInt == 0 && UPInt == 0 && name == "" {
		err := pdb.DB.Where("category_id = ?", category).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt == 0 && name == "" {
		err := pdb.DB.Where("category_id = ?", category).
			Where("price <= ?", uint(UPInt)).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if UPInt == 0 && name == "" {
		err := pdb.DB.Where("category_id = ?", category).
			Where("price >= ?", uint(LPInt)).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt != 0 && UPInt != 0 && name == "" {
		err := pdb.DB.Where("category_id = ?", category).Where("price >= ?", uint(LPInt)).
			Where("price <= ?", uint(UPInt)).Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt == 0 && UPInt == 0 && name != "" {
		err := pdb.DB.Where("category_id = ?", category).
			Where("title LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if LPInt == 0 && name != "" {
		err := pdb.DB.Where("category_id = ?", category).
			Where("price <= ?", uint(UPInt)).
			Where("title LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else if UPInt == 0 && name != "" {
		err := pdb.DB.Where("category_id = ?", category).
			Where("price >= ?", uint(LPInt)).
			Where("title LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		err := pdb.DB.Where("category_id = ?", category).Where("price >= ?", uint(LPInt)).
			Where("price <= ?", uint(UPInt)).
			Where("title LIKE ?", "%"+name+"%").Find(&products).Error
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return products, nil
}

// CreateSeller creates a new Seller in the DB
func (pdb *PostgresDb) CreateSeller(user *models.Seller) (*models.Seller, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = true
	err = pdb.DB.Create(user).Error
	return user, err
}

// CreateBuyer creates a new Buyer in the DB
func (pdb *PostgresDb) CreateBuyer(user *models.Buyer) (*models.Buyer, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = true
	err = pdb.DB.Create(user).Error
	return user, err
}

//CreateBuyerCart creates a new cart for the buyer
func (pdb *PostgresDb) CreateBuyerCart(cart *models.Cart) (*models.Cart, error) {
	var err error
	cart.CreatedAt = time.Now()
	err = pdb.DB.Create(cart).Error
	return cart, err
}

// FindSellerByUsername finds a user by the username
func (pdb *PostgresDb) FindSellerByUsername(username string) (*models.Seller, error) {
	user := &models.Seller{}

	if err := pdb.DB.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	if !user.IsActive {
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
	if !buyer.IsActive {
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
func (pdb *PostgresDb) BuyerUpdatePassword(password, newPassword string) (*models.Buyer, error) {
	buyer := &models.Buyer{}
	if err := pdb.DB.Model(buyer).Where("password_hash =?", password).Update("password_hash", newPassword).Error; err != nil {
		return nil, err
	}
	return buyer, nil
}
func (pdb *PostgresDb) SellerUpdatePassword(password, newPassword string) (*models.Seller, error) {
	seller := &models.Seller{}
	if err := pdb.DB.Model(seller).Where("password_hash =?", password).Update("password_hash", newPassword).Error; err != nil {
		return nil, err
	}
	return seller, nil
}
func (pdb *PostgresDb) BuyerResetPassword(email, newPassword string) (*models.Buyer, error) {
	buyer := &models.Buyer{}
	if err := pdb.DB.Model(buyer).Where("email =?", email).Update("password_hash", newPassword).Error; err != nil {
		return nil, err
	}
	return buyer, nil
}
