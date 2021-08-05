package routers

import (
	"apitest/database"
	"apitest/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}

	//r.GET("", api.GetBooks)
	r.GET("/:id", api.GetBook)
	//r.POST("", api.CreateBook)
	//r.PUT("/:id", api.UpdateBook)
	//r.DELETE("/:id", api.DeleteBook)

	return r
}