package test

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	mock_database "github.com/decadevs/shoparena/database/mocks"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/router"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHandleGetSellerShopByProfileAndProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	h := &handlers.Handler{
		DB: mockDB,
	}

	route, _ := router.SetupRouter(h)

	//Declaring the FAKE testing variables
	sellerFirstName := gofakeit.FirstName()
	sellerLastName := gofakeit.LastName()
	sellerEmail := gofakeit.Email()
	sellerUserName := gofakeit.Username()
	sellerPhone := gofakeit.Phone()
	sellerImage := gofakeit.ImageURL(200, 500)
	sellerStatus := gofakeit.Bool()

	productName := gofakeit.CarModel()
	productPrice := gofakeit.Price(2500, 5000)
	price := uint(productPrice)
	productCat := gofakeit.CarModel()
	sellerID := uint(gofakeit.Number(0, 10))
	testShopID := strconv.Itoa(int(sellerID))
	token := ""
	productImage := gofakeit.ImageURL(200, 500)

	//instantiating the gorm model object/struct
	testGormModel := gorm.Model{
		ID:        sellerID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	//product category object
	category := models.Category{
		testGormModel,
		productCat,
	}

	//image object
	image := models.Image{
		testGormModel,
		sellerID,
		productImage,
	}

	sliceOfImg := []models.Image{image}

	//instantiating the user model object/struct
	testUser := models.User{
		testGormModel,
		sellerFirstName,
		sellerLastName,
		sellerEmail,
		sellerUserName,
		"",
		"",
		sellerPhone,
		"",
		"",
		sellerImage,
		sellerStatus,
		token,
	}

	products := []models.Product{
		{testGormModel,
			sellerID,
			sellerID,
			category,
			productCat,
			productName,
			price,
			sliceOfImg,
			uint(rand.Intn(20)),
			sellerID,
		},
		{testGormModel,
			sellerID,
			sellerID,
			category,
			productCat,
			"",
			price,
			sliceOfImg,
			uint(rand.Intn(5)),
			sellerID,
		},
	}

	//instantiating the seller model object/struct
	testSeller := models.Seller{
		User:      testUser,
		Product:   products,
		Rating:    string(rune((rand.Intn(5)))),
		ProductId: sellerID,
	}

	bodyJSON, err := json.Marshal(testSeller)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for Error", func(t *testing.T) {
		mockDB.EXPECT().FindIndividualSellerShop(testShopID).Return(nil, errors.New("Error Exist"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/seller/"+strconv.Itoa(int(sellerID)), strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		//fmt.Println(products)
		assert.Equal(t, http.StatusNotFound, rw.Code)
		assert.Contains(t, rw.Body.String(), "Error Exist")
	})

	t.Run("Testing for success", func(t *testing.T) {
		mockDB.EXPECT().FindIndividualSellerShop(testShopID).Return(&testSeller, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/seller/"+strconv.Itoa(int(sellerID)), strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Found Seller")
	})

}
