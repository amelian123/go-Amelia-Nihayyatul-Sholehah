package routes

import (
	"book/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetBlogsController)
	e.GET("/users/:id", controllers.GetBlogsController)
	e.POST("/users", controllers.CreateBlogsController)
	e.DELETE("/users/:id", controllers.DeleteBlogsController)
	e.PUT("/users/:id", controllers.UpdateBlogsController)

	return e
}
