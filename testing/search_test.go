package testing

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/decadevs/shoparena/database"
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
	mockDB := database.NewMockDB(ctrl)
	h := &handlers.Handler{
		DB: mockDB,
	}
	route, _ := router.SetupRouter(h)

	//testString := "apple"
	product := []models.Product{
		{
			ShopName:        "the",
			ProductName:     "apple",
			ProductPrice:    "15",
			ProductCategory: "fruit",
			ProductImage:    "5678",
			ProductDetails:  "ewregfhjzn",
			Rating:          "4",
			Quantity:        "2",
		},
		{
			ShopName:        "chuk",
			ProductName:     "carburetor",
			ProductPrice:    "4000",
			ProductCategory: "spare parts",
			ProductImage:    "7685",
			ProductDetails:  "car parts",
			Rating:          "2",
			Quantity:        "20",
		},
	}

	bodyJSON, err := json.Marshal(product)
	if err != nil {
		t.Fail()
	}
	t.Run("Testing for error", func(t *testing.T) {
		mockDB.EXPECT().SearchDB(gomock.Any()).Return(nil, errors.New("error exists"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/searchproducts", strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "error exists")
	})
	t.Run("Testing for successful", func(t *testing.T) {
		mockDB.EXPECT().SearchDB(gomock.Any()).Return(product, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/searchproducts", strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(rw, req)
		fmt.Println(rw.Body.String())
		fmt.Println(product)
		assert.Equal(t, http.StatusFound, rw.Code)
		assert.Contains(t, rw.Body.String(), string(bodyJSON))
	})
}
