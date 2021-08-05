package apitest

import (
	"apitest/database"
	"apitest/handlers"
	"apitest/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetBook_OK(t *testing.T) {
	assertInstance := assert.New(t)
	database.Setup()
	db := database.GetDB()

	book, err := insertTestBook(db)
	if err != nil {
		assertInstance.Error(err)
	}

	req, w := setGetBookRouter(db, "/1")

	assertInstance.Equal(http.MethodGet, req.Method, "HTTP request method error")
	assertInstance.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		assertInstance.Error(err)
	}

	actual := models.Book{}
	if err := json.Unmarshal(body, &actual); err != nil {
		assertInstance.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := book
	expected.Model = gorm.Model{}
	fmt.Println(expected,actual)
	assertInstance.Equal(expected, actual)
	database.ClearTable()
	defer db.Close()
}

func setGetBookRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder) {
	router := gin.New()
	api := &handlers.APIEnv{DB: db}
	router.GET("/:id", api.GetBook)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return req, recorder
}

func insertTestBook(db *gorm.DB) (models.Book, error) {
	book := models.Book{
		Author:    "test",
		Name:      "test",
		PageCount: 10,
	}

	if err := db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}