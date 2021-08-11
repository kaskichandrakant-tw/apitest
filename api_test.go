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

//type IntegrationTestSuit struct {
//	suite.Suite
//}

func Test_get_book_with_id_returns_book_with_status_code_200(t *testing.T) {
	assertInstance := assert.New(t)
	dbConnection, db := getDb()

	defer dbConnection.ClearTable()
	defer db.Close()

	book, err := seedTestData(db)

	if err != nil {
		assertInstance.Error(err)
	}

	router := setupRouter(db)
	req,w :=makeApiRequest(router,"/book/1")

	assertInstance.Equal(http.MethodGet, req.Method,)
	assertInstance.Equal(http.StatusOK, w.Code, )

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
}

func getDb() (database.DbConnection ,*gorm.DB) {
	connection := database.DbConnection{}
	connection.Setup()
	db := connection.GetDB()
	return connection,db
}

func setupRouter(db *gorm.DB) (router *gin.Engine){
	router = gin.New()
	api := &handlers.APIEnv{DB: db}
	router.GET("/book/:id", api.GetBook)
	return router
}

func makeApiRequest(router *gin.Engine,url string) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	return req, recorder
}

func seedTestData(db *gorm.DB) (models.Book, error) {
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