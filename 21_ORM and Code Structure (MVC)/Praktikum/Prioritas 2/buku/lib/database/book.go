package database

import (
	"book/config"
	"book/models"
)

func GetBook() (interface{}, error) {
	var book []models.Book

	if e := config.DB.Find(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}
