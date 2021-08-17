package database

import (
	"apitest/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type DbConnection struct {
	DB *gorm.DB
}

func (connection *DbConnection) Setup() {
	if connection.DB != nil {
		return
	}
	host := "localhost"
	port := "5432"
	dbname := "test"
	user := "postgres"
	password := "password"

	db, err := gorm.Open("postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbname+
			" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate([]models.Book{})
	connection.DB = db
}

func (connection *DbConnection) GetDB() *gorm.DB {
	return connection.DB
}

func (connection *DbConnection) ClearTable() {
	connection.DB.Exec("DELETE FROM books")
	connection.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}
