package testing

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/decadevs/shoparena/database"
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/routers"
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
	route, _ := routers.SetupRouter(h)

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
		//assert.Contains(t, rw.Body.String(), product)
	})
}

//var (
//	router *gin.Engine
//)
//
//type TestHandler struct {
//	DB database.DB
//}
//
//func (t *TestHandler) SearchFunc(c *gin.Context) {
//	products := t.DB.SearchDB(c)
//	c.JSON(http.StatusFound, products)
//}
//
//func TestMain(m *testing.M) {
//	host := "localhost"
//	user := "postgres"
//	password := "donotshare"
//	dbName := "oja"
//	port := "5432"
//
//	DB := &database.PostgresDb{}
//	h := &handlers.Handler{DB: DB}
//
//	err := DB.Init(host, user, password, dbName, port)
//	if err != nil {
//		panic(err)
//	}
//
//	router, _ = routers.SetupRouter(h)
//
//	exitCode := m.Run()
//	os.Exit(exitCode)
//}
//
//// send a request to a route and get back the response
//func newRequest(method, path, token string, body interface{}) (code int, data gin.H, err error) {
//	jsonBody, err := json.Marshal(body)
//	if err != nil {
//		return 0, nil, err
//	}
//
//	w := httptest.NewRecorder()
//	req, err := http.NewRequest(method, path, strings.NewReader(string(jsonBody)))
//	if err != nil {
//		return 0, nil, err
//	}
//	req.Header.Add("Content-Type", "application/json")
//	if token != "" {
//		req.Header.Add("Authorization", "Bearer "+token)
//	}
//
//	//why do I need this?
//	router.ServeHTTP(w, req)
//
//	var response gin.H
//	err = json.Unmarshal(w.Body.Bytes(), &response)
//	return w.Code, response, err
//}
//
//func TestSearch(t *testing.T) {
//	handler := &TestHandler{}
//
//	// send the request
//
//	type TestCase struct {
//		Name            string
//		Body            string
//		ExpectedCode    int
//		ExpectedMessage string
//	}
//
//	testString := "chucks"
//	testcases := []TestCase{
//		{
//			Name:         "valid search",
//			Body:         testString,
//			ExpectedCode: http.StatusFound,
//		},
//	}
//
//	for _, testcase := range testcases {
//
//		code, _, err := newRequest(http.MethodGet, "/products/searchproducts", "", testcase.Body)
//		if err != nil {
//			t.Errorf("%s: error sending request: %v", testcase.Name, err)
//		}
//
//		handler.SearchFunc()
//		// assert the Code
//		if code != testcase.ExpectedCode {
//			t.Errorf("%s: expected status code %d, got %d", testcase.Name, testcase.ExpectedCode, code)
//		}
//
//	}
//	//
//	//if err := db.Delete(&models.Product{}, "string = ?", testcases[0].Body.ProductName).Error; err != nil {
//	//	t.Error("Error deleting user")
//	//}
//}
