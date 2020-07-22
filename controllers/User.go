package controllers

import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/repository"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var userx []models.User
	err := repository.GetAllUsers(&userx)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound,)
	} else {
		c.JSON(http.StatusOK, userx)
	}
}

//CreateUser ... Create user
func CreateUser(c *gin.Context) {
	var userx models.User

	if err := c.ShouldBindJSON(&userx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&userx)
	err := repository.CreateUser(&userx)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userx)
	}
}

//GetSearchUser ... Get the user by name
func GetSearchUser(c *gin.Context) {
	name := c.Params.ByName("name") 
	var userx models.User
	err := repository.GetSearchUser(&userx, name)
	if err != nil {
		errorx:=helpers.Errorx{Msgx:"There is no such record",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, userx)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var userx models.User
	id := c.Params.ByName("id")
	err := repository.GetSearchUser(&userx, id) 
	if err != nil {
		c.JSON(http.StatusNotFound, userx)
	}
	c.BindJSON(&userx)
	err = repository.UpdateUser(&userx, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userx)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var userx models.User
	id := c.Params.ByName("id")
	err := repository.DeleteUser(&userx, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}