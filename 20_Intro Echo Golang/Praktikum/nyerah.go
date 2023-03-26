package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// get user id from path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user id",
		})
	}

	// find user by id
	var user User
	for _, u := range users {
		if u.Id == id {
			user = u
			break
		}
	}

	if user.Id == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"user":     user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// get user id from path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user id",
		})
	}

	// find user by id and remove from slice
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user by id",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// get user id from path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid user id",
		})
	}

	// find user by id and update
	var updatedUser User
	for i, u := range users {
		if u.Id == id {
			// binding data
			c.Bind(&updatedUser)

			// update user
			updatedUser.Id = id
			users[i] = updatedUser

			break
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id",
		"user":     updatedUser,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.GET("/users/:id", GetUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
