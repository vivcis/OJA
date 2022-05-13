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
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	h := &handlers.Handler{
		DB: mockDB,
	}
	route, _ := router.SetupRouter(h)

	category := models.Category{
		Model: gorm.Model{ID: 1},
		Name:  "Shirt",
	}

	product := models.Product{
		Model:       gorm.Model{ID: 4, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		SellerId:    2,
		CategoryId:  category.ID,
		Title:       "household",
		Description: "home items",
		Price:       78,
		Rating:      8,
		Quantity:    4903,
	}

	bodyJSON, err := json.Marshal(product)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for valid update", func(t *testing.T) {
		mockDB.EXPECT().UpdateProductByID(product).Return(nil)
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPut,
			"/api/v1/update/product/4",
			strings.NewReader(string(bodyJSON)))
		if err != nil {
			fmt.Printf("errrr here %v \n", err)
			return
		}
		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
	})
}
