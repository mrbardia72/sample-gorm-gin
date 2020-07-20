package controllers

import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/repository"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetProfiles ... Get all profiles
func GetProfiles(c *gin.Context) {
	var profilex []models.Profile
	err := repository.GetAllProfiles(&profilex)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profilex)
	}
}

//CreateProfile ... Create Profile
func CreateProfile(c *gin.Context) {
	var profilex models.Profile

	if err := c.ShouldBindJSON(&profilex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&profilex)
	err := repository.CreateProfile(&profilex)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profilex)
	}
}

//GetSearch ... Get the user by name
func GetSearch(c *gin.Context) {
	name := c.Params.ByName("name") 
	var profilex models.Profile
	err := repository.GetSearch(&profilex, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profilex)
	}
}

//Updateprofile ... Update the profile information
func UpdateProfile(c *gin.Context) {
	var profilex models.Profile
	id := c.Params.ByName("id")
	err := repository.GetProfileByID(&profilex, id)
	if err != nil {
		c.JSON(http.StatusNotFound, profilex)
	}
	c.BindJSON(&profilex)
	err = repository.UpdateProfile(&profilex, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profilex)
	}
}

//Deleteprofile ... Delete the profile
func DeleteProfile(c *gin.Context) {
	var profilex models.Profile
	id := c.Params.ByName("id")
	err := repository.DeleteProfile(&profilex, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
