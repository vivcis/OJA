package test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	mock_database "github.com/decadevs/shoparena/database/mocks"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/router"
	"github.com/decadevs/shoparena/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetBuyerDetailsHandler(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println(err.Error())
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_database.NewMockDB(ctrl)
	h := &handlers.Handler{DB: mockDB}
	route, _ := router.SetupRouter(h)

	accClaim, _ := services.GenerateClaims("cecilia@yahoo.com")

	secret := os.Getenv("JWT_SECRET")
	acc, err := services.GenerateToken(jwt.SigningMethodHS256, accClaim, &secret)
	if err != nil {
		t.Fail()
	}

	getBuyer := &models.User{
		FirstName:   "Cecilia",
		LastName:    "Orji",
		Address:     "Edo Tech Park",
		PhoneNumber: "09052755438",
		Email:       "cecilia@yahoo.com",
		Username:    "cece",
		Image:       "",
		Token:       "",
        IsActive:    true,

	}
	buyer := models.Buyer{
		User: models.User{
			Model: gorm.Model{
				ID: 23,
				CreatedAt: time.Time{},
		        UpdatedAt: time.Time{},
			},

			FirstName:   "Cecilia",
			LastName:    "Orji",
			PhoneNumber: "09052755438",
			Email:       "cecilia@yahoo.com",
			Address:     "Edo Tech Park",
		},
		Orders: nil,
	}
	testBuyer, err := json.Marshal(getBuyer)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	mockDB.EXPECT().TokenInBlacklist(gomock.Any()).Return(false).Times(2)
	mockDB.EXPECT().FindBuyerByEmail(buyer.Email).Return(&buyer, nil).Times(2)

	t.Run("Test for error", func(t *testing.T) {
		//mockDB.EXPECT().UpdateBuyerProfile(buyer.ID, getBuyer).Return(errors.New("error exist"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/getbuyerprofile",
			strings.NewReader(string(testBuyer)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))
		route.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "internal server error")
	})

	t.Run("Test for successful retrieval", func(t *testing.T) {
		//mockDB.EXPECT().UpdateBuyerProfile(buyer.ID, getBuyer).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/getbuyerprofile",
			strings.NewReader(string(testBuyer)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))
		route.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "buyer details retrieved correctly")
	})

}