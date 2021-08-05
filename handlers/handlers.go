package handlers

import (
	"apitest/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

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

//func (a *APIEnv) GetBooks(c *gin.Context) {
//	books, err := database.GetBooks(a.DB)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	c.JSON(http.StatusOK, books)
//}



//func (a *APIEnv) CreateBook(c *gin.Context) {
//	book := models.Book{}
//	err := c.BindJSON(&book)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	if err := a.DB.Create(&book).Error; err != nil {
//		c.JSON(http.StatusInternalServerError,err.Error())
//		return
//	}
//
//	c.JSON(http.StatusOK, book)
//}
//
//func (a *APIEnv)DeleteBook(c *gin.Context) {
//	id := c.Params.ByName("id")
//	_, exists, err := database.GetBookByID(id,a.DB)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	if !exists {
//		c.JSON(http.StatusNotFound, "record not exists")
//		return
//	}
//
//	err = database.DeleteBook(id,a.DB)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	c.JSON(http.StatusOK, "record deleted successfully")
//}
//
//func (a *APIEnv) UpdateBook(c *gin.Context) {
//	id := c.Params.ByName("id")
//	_, exists, err := database.GetBookByID(id, a.DB)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	if !exists {
//		c.JSON(http.StatusNotFound, "record not exists")
//		return
//	}
//
//	updatedBook := models.Book{}
//	err = c.BindJSON(&updatedBook)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	if err := database.UpdateBook(a.DB, &updatedBook); err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	a.GetBook(c)
//}