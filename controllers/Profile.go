package controllers

import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/models/profile"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetProfiles ... Get all profiles
func GetProfiles(c *gin.Context) {
	var profilex []profile.Profile
	err := profile.GetAllProfiles(&profilex)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profilex)
	}
}

//CreateProfile ... Create Profile
func CreateProfile(c *gin.Context) {

	var profilex profile.Profile
	c.BindJSON(&profilex)
	err := profile.CreateProfile(&profilex)
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
	var profilex profile.Profile
	err := profile.GetSearch(&profilex, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profilex)
	}
}

//Updateprofile ... Update the profile information
func UpdateProfile(c *gin.Context) {
	var profilex profile.Profile
	id := c.Params.ByName("id")
	err := profile.GetProfileByID(&profilex, id)
	if err != nil {
		c.JSON(http.StatusNotFound, profilex)
	}
	c.BindJSON(&profilex)
	err = profile.UpdateProfile(&profilex, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profilex)
	}
}

//Deleteprofile ... Delete the profile
func DeleteProfile(c *gin.Context) {
	var profilex profile.Profile
	id := c.Params.ByName("id")
	err := profile.DeleteProfile(&profilex, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
