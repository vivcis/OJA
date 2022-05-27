
//package handlers
//
//import (
//	"fmt"
//	"github.com/decadevs/shoparena/models"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"time"
//)
//
//func (h *Handler) CreateProducts(c *gin.Context) {
//	user, exist := c.Get("user")
//	if !exist {
//		c.JSON(http.StatusInternalServerError, "not logged in")
//	}
//
//	seller := user.(*models.Seller)
//
//	product := models.Product{}
//
//	err := c.BindJSON(&product)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"error": "Unable to bind product",
//		})
//		return
//	}
//
//	product.SellerId = seller.ID
//	fmt.Println(product.SellerId)
//
//	err = h.DB.CreateProduct(product)
//	if err != nil {
//		log.Println("check error: ", err)
//		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating product in handler"})
//		return
//	}
//
//	response := gin.H{
//
//		"data":      product,
//		"status":    http.StatusText(http.StatusCreated),
//		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
//	}
//
//	c.JSON(http.StatusCreated, response)
//}


package handlers

import (
	"fmt"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/server/response"
	"github.com/decadevs/shoparena/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func (h *Handler) CreateProducts(c *gin.Context) {
	//get the user id from a logged-in user
	user, exists := c.Get("user")
	if !exists {
		log.Printf("can't get user from context\n")
		response.JSON(c, "", http.StatusUnauthorized, nil, []string{"you are not logged in"})
		return
	}
	seller := user.(*models.Seller)
	//userId := userI.(*models.User).ID

	form, err := c.MultipartForm()

	if err != nil {
		log.Printf("error parsing multipart form: %v", err)
		response.JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	formImages := form.File["images"]
	images := []models.Image{}

	// upload the images to aws.
	for _, f := range formImages {
		file, err := f.Open()
		if err != nil {

		}
		fileExtension, ok := services.CheckSupportedFile(strings.ToLower(f.Filename))
		log.Printf(filepath.Ext(strings.ToLower(f.Filename)))
		fmt.Println(fileExtension)
		if ok {
			log.Println(fileExtension)
			response.JSON(c, "", http.StatusBadRequest, nil, []string{fileExtension + " image file type is not supported"})
			return
		}

		session, tempFileName, err := services.PreAWS(fileExtension, "product")
		if err != nil {
			log.Println("could not upload file", err)
		}

		url, err := h.DB.UploadFileToS3(session, file, tempFileName, f.Size)
		if err != nil {
			log.Println(err)
			response.JSON(c, "", http.StatusInternalServerError, nil, []string{"an error occurred while uploading the image"})
			return
		}

		log.Printf("filename: %v", f.Filename)

		img := models.Image{
			Url: url,
		}
		images = append(images, img)
	}

	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		response.JSON(c, "", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	// err := strconv.Atoi(c.PostForm("no_of_rooms"))
	rating, err := strconv.Atoi(c.PostForm("rating"))
	if err != nil {
		response.JSON(c, "", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		response.JSON(c, "", http.StatusBadRequest, nil, []string{err.Error()})
	}

	products := models.Product{
		Model: gorm.Model{
			ID: 1,
		},
		Category: models.Category{
			Name: c.PostForm("name"),
			Model: gorm.Model{
				ID: 1,
			},
		},
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Price:       uint(price),
		SellerId:    seller.ID,
		Images:      images,
		Rating:      uint(rating),
		Quantity:    uint(quantity),
	}

	err = h.DB.CreateProduct(products)
	if err != nil {
		response.JSON(c, "", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	log.Println(products.Images)
	response.JSON(c, "Product Successfully Added", http.StatusOK, products, nil)

=======
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) CreateProducts(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, "not logged in")
	}

	seller := user.(*models.Seller)

	product := models.Product{}

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to bind product",
		})
		return
	}

	product.SellerId = seller.ID
	fmt.Println(product.SellerId)

	err = h.DB.CreateProduct(product)
	if err != nil {
		log.Println("check error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating product in handler"})
		return
	}

	response := gin.H{

		"data":      product,
		"status":    http.StatusText(http.StatusCreated),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, response)
>>>>>>> 756bc9c4b946eb4d4cfc2fa626fee6ea258d66d9
}
