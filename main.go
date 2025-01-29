package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct to represent data
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Mock database
var users = []User{
	{ID: 1, Name: "priya", Age: 25},
	{ID: 2, Name: "Nethu", Age: 30},
}

func main() {
	router := gin.Default()

	// Define API routes
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)

	// Start server on port 8080
	router.Run(":8080")
}

// Handler to get all users
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// Handler to get user by ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if id == string(rune(user.ID)) {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// Handler to create a new user
func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}
