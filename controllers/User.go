package controllers

import (
	// "fmt"
	repo_user "github.com/mrbardia72/sample-gorm-gin/repository/user"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
  
// //GetUsers ... Get all users
// func GetUsers(c *gin.Context) {
// 	var userx []models.User
// 	err := repo_user.GetAllUsers(&userx)

// 	if err != nil {
// 		errorx:=helpers.Errorx{Msgx:"error not get all users becaus empty table user",Codex:"404"}
// 		c.JSON(http.StatusNotFound,errorx)
// 	} else {
// 		c.JSON(http.StatusOK, userx)
// 	}

// }

// //CreateUser ... Create user
// func CreateUser(c *gin.Context) {
// 	var userx models.User

// 	if err := c.ShouldBindJSON(&userx); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.BindJSON(&userx)
// 	err := repo_user.CreateUser(&userx)

// 	if err != nil {
// 		errorx:=helpers.Errorx{Msgx:"error create to db",Codex:"404"}
// 		c.JSON(http.StatusNotFound,errorx)
// 	} else {
// 		c.JSON(http.StatusOK, userx)
// 	}
// }

// //GetSearchUser ... Get the user by name
// func GetSearchUser(c *gin.Context) {
// 	name := c.Params.ByName("name") 
// 	var userx models.User
// 	err := repo_user.GetSearchUser(&userx, name)
// 	if err != nil {
// 		errorx:=helpers.Errorx{Msgx:"There is no such recordxxx",Codex:"404"}
// 		c.JSON(http.StatusNotFound,errorx)
// 	} else {
// 		c.JSON(http.StatusOK, userx)
// 	}
// }

//UpdateUser ... Update the user information
// func UpdateUser(c *gin.Context) {
// 	var userx models.User
// 	id := c.Params.ByName("id")
// 	err := repo_user.GetSearchUser(&userx, id) 
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, userx)
// 	}
// 	c.BindJSON(&userx)
// 	err = repo_user.UpdateUser(&userx, id)
// 	if err != nil {
// 		errorx:=helpers.Errorx{Msgx:"not update a record",Codex:"404"}
// 		c.JSON(http.StatusNotFound,errorx)
// 	} else {
// 		c.JSON(http.StatusOK, userx)
// 	}
// }

// //DeleteUser ... Delete the user
// func DeleteUser(c *gin.Context) {
// 	var userx models.User
// 	id := c.Params.ByName("id")
// 	err := repo_user.DeleteUser(&userx, id)

// 	if err != nil {
// 		errorx:=helpers.Errorx{Msgx:"not delete a record",Codex:"404"}
// 		c.JSON(http.StatusNotFound,errorx)
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
// 	}
	
// }