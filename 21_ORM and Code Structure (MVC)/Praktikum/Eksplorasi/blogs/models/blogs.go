package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Kategori string `json:"kategori" form:"kategori"`
	Penulis  string `json:"penulis" form:"penulis"`
}
