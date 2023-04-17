package controllers

import (
	"book/config"
	"book/models"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// get all book
func GetBookController(c echo.Context) error {
	var book []models.Books

	if err := config.DB.Find(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all book",
		"book":    book,
	})
}

// get book by id
func GetbookController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var book models.Books
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "berhasil",
		"data_book": book,
	})

}

// create new book
func CreatebookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var book models.Book
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "book not available",
		})
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal hapus data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil menghapus data",
		"book":   book,
	})
}

// update book by id
func UpdatebookController(c echo.Context) error {
	// your solution here
	body := new(models.Jilid)

	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameter",
		})
	}

	var book models.Jilid
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "book not available",
		})
	}

	book.Judul = body.Judul
	book.Penerbit = body.Penerbit
	book.Penulis = body.Penulis

	if err := config.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil update data",
		"book":   book,
	})
}
