package routers

import (
	"apitest/database"
	"apitest/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	databaseConnection := database.DbConnection{}
	databaseConnection.Setup()

	api := &handlers.APIEnv{
		DB: databaseConnection.GetDB(),
	}
	r.GET("/book/:id", api.GetBook)
	r.GET("/usio/bank", api.GetUSIOBank)
	return r
}
