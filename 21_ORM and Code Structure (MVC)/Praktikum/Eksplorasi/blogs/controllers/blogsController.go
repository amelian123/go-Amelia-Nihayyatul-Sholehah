package controllers

import (
	"book/config"
	"book/models"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.blogss

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blogs by id
func GetBlogsController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var blogs models.blogs
	if err := config.DB.First(&blogs, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "blogs not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "berhasil",
		"data_blogs": blogs,
	})

}

// create new blogs
func CreateBlogsController(c echo.Context) error {
	blogs := models.blogs{}
	c.Bind(&blogs)

	if err := config.DB.Save(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blogs",
		"blogs":   blogs,
	})
}

// delete blogs by id
func DeleteBlogsController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var blogs models.blogs
	if err := config.DB.Where("id = ?", id).First(&blogs).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "blogs not available",
		})
	}

	if err := config.DB.Delete(&blogs).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal hapus data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil menghapus data",
		"blogs":  blogs,
	})
}

// update blogs by id
func UpdateBlogsController(c echo.Context) error {
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

	var blogs models.Jilid
	if err := config.DB.Where("id = ?", id).First(&blogs).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "blogs not available",
		})
	}

	blogs.Judul = body.Judul
	blogs.Kategori = body.Kategori
	blogs.Penulis = body.Penulis

	if err := config.DB.Save(&blogs).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil update data",
		"blogs":  blogs,
	})
}
