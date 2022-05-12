package testing

import (
	"testing"

	mock_database "github.com/decadevs/shoparena/database/mocks"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/router"
	"github.com/golang/mock/gomock"
)

func TestUpdateProfile(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	h := &handlers.Handler{
		//DB: mockDB,
	}
	route, _ := router.SetupRouter(h)
}