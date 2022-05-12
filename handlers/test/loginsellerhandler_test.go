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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSellerLogin(t *testing.T) {

	//import new controller
	ctxl := gomock.NewController(t)
	defer ctxl.Finish()

	mockdb := mock_database.NewMockDB(ctxl)

	mockhandler := &handlers.Handler{
		DB: mockdb,
	}

	rout, _ := router.SetupRouter(mockhandler)
	t.Run("testing bad request", func(t *testing.T) {
		usermk := &struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Username: "",
			Password: "12345566666",
		}
		bytes, _ := json.Marshal(usermk)
		ht := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/loginseller", strings.NewReader(string(bytes)))

		rout.ServeHTTP(ht, req)
		assert.Equal(t, http.StatusBadRequest, ht.Code)
		assert.Contains(t, ht.Body.String(), "bad request")
	})
	t.Run("find seller by username", func(t *testing.T) {
		hash, _ := handlers.HashPassword("12345566666")
		seller := &models.Seller{
			User: models.User{
				Username:     "mike123",
				PasswordHash: hash,
			},
		}
		usermk := &struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Username: "mike123",
			Password: "12345566666",
		}
		mockdb.EXPECT().FindSellerByUsername(seller.Username).Return(seller, nil)
		bytes, _ := json.Marshal(usermk)
		ht := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/loginseller", strings.NewReader(string(bytes)))

		rout.ServeHTTP(ht, req)
		assert.Equal(t, http.StatusOK, ht.Code)
		fmt.Println(ht.Body.String())
		assert.Contains(t, ht.Body.String(), usermk.Username)
	})

}
