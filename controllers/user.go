package controllers

import (
	"api-server/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers fetches all user data.
func GetAllUsers(c *gin.Context) {
	var user []models.User

	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser inserts a new user.
func CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUserByID fetches a user by ID.
func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser updates a user.
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)

	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser deletes a user.
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted."})
	}
}
