package database

import (
	"apitest/models"
	"github.com/jinzhu/gorm"
)

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