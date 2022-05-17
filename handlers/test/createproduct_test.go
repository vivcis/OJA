package test

import (
	"encoding/json"
	"fmt"
	mock_database "github.com/decadevs/shoparena/database/mocks"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/router"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"

	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	h := &handlers.Handler{
		DB: mockDB,
	}
	route, _ := router.SetupRouter(h)

	product := models.Product{

		SellerId:    1,
		CategoryId:  2,
		Title:       "plumbing",
		Description: "Building",
		Price:       6700,
		Rating:      5,
		Quantity:    3,
	}

	bodyJSON, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	mockDB.EXPECT()

	t.Run("testing for created update", func(t *testing.T) {
		mockDB.EXPECT().CreateProduct(product).Return(nil)
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/api/v1/createproduct/1",
			strings.NewReader(string(bodyJSON)))
		if err != nil {
			fmt.Printf("error occured %v \n", err)
			return
		}
		t.Log(err)
		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusCreated, rw.Code)
		assert.Contains(t, rw.Body.String(), product.Description)
	})
}
