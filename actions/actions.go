package actions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// User represents a user entity
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Mock database
var users = map[string]User{
	"1": {ID: "1", Name: "John Doe"},
	"2": {ID: "2", Name: "Jane Doe"},
}

// GetUsers returns a list of all users
func GetUsers(c echo.Context) error {
	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}
	return c.JSON(http.StatusOK, userList)
}

// GetUser returns a single user by ID
func GetUser(c echo.Context) error {
	id := c.Param("id")
	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser adds a new user
func CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	users[user.ID] = *user
	return c.JSON(http.StatusCreated, user)
}

// UpdateUser updates an existing user
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	if _, exists := users[id]; !exists {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	users[id] = *user
	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if _, exists := users[id]; !exists {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	delete(users, id)
	return c.JSON(http.StatusOK, "User deleted")
}
