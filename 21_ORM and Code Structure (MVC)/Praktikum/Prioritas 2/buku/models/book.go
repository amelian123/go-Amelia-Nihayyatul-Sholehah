package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	Penulis  string `json:"penulis" form:"penulis"`
}
