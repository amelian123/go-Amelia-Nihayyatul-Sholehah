package database

import (
	"book/config"
)

func Getblogs() (interface{}, error) {
	var blogs []models.blogs

	if e := config.DB.Find(&blogs).Error; e != nil {
		return nil, e
	}
	return blogs, nil
}
