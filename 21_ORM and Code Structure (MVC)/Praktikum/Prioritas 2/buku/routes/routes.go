package routes

import (
	"book/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetBookController)
	e.GET("/users/:id", controllers.GetBookController)
	e.POST("/users", controllers.CreateBookController)
	e.DELETE("/users/:id", controllers.DeleteBookController)
	e.PUT("/users/:id", controllers.UpdateBookController)

	return e
}
