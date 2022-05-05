### User struct
* first name
* last name
* address
* username
* email
* password
* passwordhash
* Phone number
* User Image



### Buyer Struct
* User
* Cart



### Seller Struct
* User
* Rating
* Shop

### Shop Struct
* Slice of product
* Name of shop
### Basket
* Buyer username
* Product
### Product Struct
* Category
* Rating
* Product Name
* Product Image
* Product details
* Price
* Quantity

### Cart Struct
* Slice of product
* Total price
* Total quantity


# interfaces that might be needed
check postgres.go for the related db functions

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