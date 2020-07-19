package controllers

import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetProfiles ... Get all profiles
func GetProfiles(c *gin.Context) {
	var profile []models.Profile
	err := models.GetAllProfiles(&profile)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

//CreateProfile ... Create Profile
func CreateProfile(c *gin.Context) {

	var profile models.Profile
	c.BindJSON(&profile)
	err := models.CreateProfile(&profile)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

//GetSearch ... Get the user by name
func GetSearch(c *gin.Context) {
	name := c.Params.ByName("name") 
	var profile models.Profile
	err := models.GetSearch(&profile, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

//Updateprofile ... Update the profile information
func UpdateProfile(c *gin.Context) {
	var profile models.Profile
	id := c.Params.ByName("id")
	err := models.GetProfileByID(&profile, id)
	if err != nil {
		c.JSON(http.StatusNotFound, profile)
	}
	c.BindJSON(&profile)
	err = models.UpdateProfile(&profile, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

//Deleteprofile ... Delete the profile
func DeleteProfile(c *gin.Context) {
	var profile models.Profile
	id := c.Params.ByName("id")
	err := models.DeleteProfile(&profile, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
