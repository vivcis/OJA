package testing

import (
	"encoding/json"
	"fmt"
	mock_database "github.com/decadevs/shoparena/database/mocks"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/router"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	h := &handlers.Handler{
		DB: mockDB,
	}
	route, _ := router.SetupRouter(h)

	product := []models.Product{
		{
			ShopName:        "the",
			ProductName:     "apple",
			ProductPrice:    15,
			ProductCategory: "fruit",
			ProductImage:    "5678",
			ProductDetails:  "ewregfhjzn",
			Rating:          4,
			Quantity:        2,
		},
		{
			ShopName:        "chuk",
			ProductName:     "carburetor",
			ProductPrice:    4000,
			ProductCategory: "spare parts",
			ProductImage:    "7685.com",
			ProductDetails:  "car parts",
			Rating:          2,
			Quantity:        20,
		},
		{
			ShopName:        "the",
			ProductName:     "corn",
			ProductPrice:    25,
			ProductCategory: "fruit",
			ProductImage:    "5678.com",
			ProductDetails:  "ewregfhjzn",
			Rating:          4,
			Quantity:        2,
		},
	}

	//s := "apple"
	p := []models.Product{
		{
			ShopName:        "the",
			ProductName:     "apple",
			ProductPrice:    15,
			ProductCategory: "fruit",
			ProductImage:    "5678",
			ProductDetails:  "ewregfhjzn",
			Rating:          4,
			Quantity:        2,
		},

		{
			ShopName:        "the",
			ProductName:     "corn",
			ProductPrice:    25,
			ProductCategory: "fruit",
			ProductImage:    "5678.com",
			ProductDetails:  "ewregfhjzn",
			Rating:          4,
			Quantity:        2,
		},
	}

	pJSON, err := json.Marshal(p)
	if err != nil {
		t.Fail()
	}

	bodyJSON, err := json.Marshal(product)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for error", func(t *testing.T) {
		mockDB.EXPECT().SearchProduct("", "", "", "")
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/searchproducts", strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "no such product")

	})

	t.Run("Testing for nonempty string", func(t *testing.T) {
		mockDB.EXPECT().SearchProduct("25", "30", "fruit", "").Return(p, nil)
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet,
			"/api/v1/searchproducts?category=fruit&lower-price=25&upper-price=30",
			strings.NewReader(string(pJSON)))
		if err != nil {
			fmt.Printf("errrr here %v \n", err)
			return
		}
		route.ServeHTTP(rw, req)
		//fmt.Println(rw.Body.String())
		assert.Equal(t, http.StatusFound, rw.Code)
		assert.Contains(t, rw.Body.String(), string(pJSON))
	})
}
