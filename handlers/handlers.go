package handlers

import (
	"apitest/database"
	"apitest/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

type ResponseLocal struct {
	Success string
	Data    map[string]string
}
type test interface {

}

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetBook(c *gin.Context) {
	id := c.Params.ByName("id")
	book, exists, err := database.GetBookByID(id, a.DB)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no book in db")
		return
	}
	c.JSON(http.StatusOK, book)
}
func (a *APIEnv) GetUSIOBank(c *gin.Context) {
	client := &http.Client{}
	usioConfig := models.USIOConfig{os.Getenv("USIO_HOST"),os.Getenv("USIO_PORT")}
	req, err := http.NewRequest("GET", usioConfig.GetUSIOUri()+"/api/v2/bank", nil)
	req.Header.Add("Authorization", c.Request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err)
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	responseData := ResponseLocal{}
	json.NewDecoder(resp.Body).Decode(&responseData)
	c.JSON(200, responseData)
}
