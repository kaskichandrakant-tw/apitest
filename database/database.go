package database

import (
"apitest/models"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/postgres"
"log"
)

/*DB is connected database object*/
var DB *gorm.DB

func Setup() {
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
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

//func GetBooks(db *gorm.DB) ([]models.Book, error) {
//	books := []models.Book{}
//	query := db.Select("books.*").
//		Group("books.id")
//	if err := query.Find(&books).Error; err != nil {
//		return books, err
//	}
//
//	return books, nil
//}

func GetBookByID(id string, db *gorm.DB) (models.Book, bool, error) {
	b := models.Book{}

	query := db.Select("books.*")
	query = query.Group("books.id")
	err := query.Where("books.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

//func DeleteBook(id string, db *gorm.DB) error {
//	var b models.Book
//	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func UpdateBook(db *gorm.DB, b *models.Book) error {
//	if err := db.Save(&b).Error; err != nil {
//		return err
//	}
//	return nil
//}

func ClearTable() {
	DB.Exec("DELETE FROM books")
	DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}
