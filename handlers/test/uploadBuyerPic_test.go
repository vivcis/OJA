package test

import (
	"fmt"
	mock_database "github.com/decadevs/shoparena/database/mocks"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/router"
	"github.com/decadevs/shoparena/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadBuyerProfielPic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	//mockMail := mock_database.NewMockMailer(ctrl)
	h := &handlers.Handler{DB: mockDB}
	route, _ := router.SetupRouter(h)

	accessClaims, _ := services.GenerateClaims("ceciliaorji@yahoo.com")
	secret := os.Getenv("JWT_SECRET")
	accToken, err := services.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		t.Fail()
	}
	file, _ := createImg()
	buyer := models.Buyer{
		User: models.User{
			Model: gorm.Model{
				ID: 23,
			},
			Username:    "cece",
			FirstName:   "Orji",
			LastName:    "Cecilia",
			PhoneNumber: "09052755438",
			Email:       "ceciliaorji.co@gmail.com",
			Address:     "Edo Tech Park",
			Image:       "https://files.fm/u/w4hjhgvex",
		},
		Orders: nil,
	}

	mockDB.EXPECT().FindBuyerByEmail(buyer.Email).Return(&buyer, nil).Times(3).AnyTimes()
	mockDB.EXPECT().TokenInBlacklist(gomock.Any()).Return(false).Times(3).AnyTimes()
	mockDB.EXPECT().UploadFileToS3(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("", nil).AnyTimes()

	t.Run("Test Upload profile pic success", func(t *testing.T) {
		b, content_type, err := prepfile(file)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/uploadbuyerpic", b)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		req.Header.Set("Content-Type", content_type)
		route.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, 200)
		assert.Contains(t, resp.Body.String(), "")
	})
}
