package test

//import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//	"time"
//
//	mock_database "github.com/decadevs/shoparena/database/mocks"
//	"github.com/decadevs/shoparena/handlers"
//	"github.com/decadevs/shoparena/models"
//	"github.com/decadevs/shoparena/router"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"gorm.io/gorm"
//)
//
//func TestGetAllSellers(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	mockDB := mock_database.NewMockDB(ctrl)
//
//	h := &handlers.Handler{DB: mockDB}
//
//	route, _ := router.SetupRouter(h)
//
//	testGorm := gorm.Model{
//		ID:        1,
//		CreatedAt: time.Time{},
//		UpdatedAt: time.Time{},
//	}
//
//	//an instance of a user
//	user := models.User{
//		testGorm,
//		"Alex",
//		"Bekaren",
//		"udu@yahoo.com",
//		"@theAlex",
//		"",
//		"",
//		"",
//		"No.3 Space Avenue, Toronto",
//		"08076453879",
//		"alex.jpg",
//		true,
//		""}
//
//	category := models.Category{
//		testGorm,
//		"Bicycles",
//	}
//	images := models.Image{
//		testGorm, 1, "gre.com",
//	}
//
//	sliceImage := []models.Image{images}
//	product := models.Product{
//		testGorm,
//		1,
//		2,
//		category,
//		"Raleigh Bicycle",
//		"steel frame, 6x speed",
//		600,
//		sliceImage,
//		4,
//		100,
//	}
//
//	sliceOfProduct := []models.Product{product}
//
//	order := models.Order{}
//	sliceOfOrder := []models.Order{order}
//	sellerID := 4
//
//	seller := models.Seller{
//		testGorm,
//		user,
//		sliceOfProduct,
//		sliceOfOrder,
//		sellerID,
//	}
//	sliceOfSeller := []models.Seller{seller}
//
//	bodyJSON, err := json.Marshal(seller)
//	if err != nil {
//		t.Fail()
//	}
//
//	t.Run("Testing For Error", func(t *testing.T) {
//		mockDB.EXPECT().GetAllSellers().Return(nil, errors.New("Error Exist"))
//		rw := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodGet, "/api/v1/sellers", strings.NewReader(string(bodyJSON)))
//		route.ServeHTTP(rw, req)
//		fmt.Println(rw.Body.String())
//		assert.Equal(t, http.StatusInternalServerError, rw.Code)
//		assert.Contains(t, rw.Body.String(), "Error Exist in")
//	})
//
//	t.Run("Testing To Get Seller Successful", func(t *testing.T) {
//		mockDB.EXPECT().GetAllSellers().Return(sliceOfSeller, nil)
//		rw := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodGet, "/api/v1/sellers", strings.NewReader(string(bodyJSON)))
//		route.ServeHTTP(rw, req)
//		fmt.Println(rw.Body.String())
//		assert.Equal(t, http.StatusOK, rw.Code)
//		assert.Contains(t, rw.Body.String(), "All Seller Found")
//	})
//
//}
