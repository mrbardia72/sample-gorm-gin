package controllers

import (
	// "fmt"
	"github.com/mrbardia72/sample-gorm-gin/repository"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

//Getposts ... Get all posts
func Getposts(c *gin.Context) {
	var postx []models.Post
	err := repository.GetAllposts(&postx)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, postx)
	}
}