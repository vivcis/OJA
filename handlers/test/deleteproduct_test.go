package test

import (
	"encoding/json"
	"fmt"
	mock_database "github.com/decadevs/shoparena/database/mocks"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/router"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Unable to load env")
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	h := &handlers.Handler{
		DB: mockDB,
	}
	route, _ := router.SetupRouter(h)

	product := models.Product{

		SellerId:    1,
		CategoryId:  1,
		Title:       "Material",
		Description: "Building",
		Price:       670,
		Rating:      10,
		Quantity:    80,
	}

	bodyJSON, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	t.Run("testing for delete product", func(t *testing.T) {
		mockDB.EXPECT().DeleteProduct(&product).Return(nil)
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, "api/v1/deleteproduct/1",
			strings.NewReader(string(bodyJSON)))
		if err != nil {
			fmt.Printf("error occured")
			return
		}
		t.Log(err)
		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusCreated, rw.Code)
		assert.Contains(t, rw.Body.String(), "product", "successfully")
	})

}
