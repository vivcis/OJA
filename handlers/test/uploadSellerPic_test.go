package test

import (
	"bytes"
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
	"image"
	"image/color"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createImg() (*os.File, error) {
	width := 200
	height := 100
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}
	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}
	// Encode as PNG.
	f, err := os.Create("testimg.png")
	if err != nil {
		return nil, err
	}
	png.Encode(f, img)
	return f, nil
}
func prepfile(file *os.File) (*bytes.Buffer, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	defer w.Close()
	fmt.Println(file.Name())
	if _, err := w.CreateFormFile("profile_picture", file.Name()); err != nil {
		return nil, "", fmt.Errorf("%v", err)
	}
	fmt.Println(w.FormDataContentType())
	return &b, w.FormDataContentType(), nil
}

func TestUploadprofielpic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)
	//mockMail := mock_database.NewMockMailer(ctrl)
	h := &handlers.Handler{DB: mockDB}
	route, _ := router.SetupRouter(h)

	accessClaims, _ := services.GenerateClaims("ceciliaorji.co@gmail.com")
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

	t.Run("Test Upload profiele pic success", func(t *testing.T) {
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
