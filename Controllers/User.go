package Controllers

import (
	"fmt"
	"../Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetProfiles ... Get all profiles
func GetProfiles(c *gin.Context) {
	var profile []Models.Profile
	err := Models.GetAllProfiles(&profile)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

//CreateProfile ... Create Profile
func CreateProfile(c *gin.Context) {

	var profile Models.Profile
	c.BindJSON(&profile)
	err := Models.CreateProfile(&profile)
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
	var profile Models.Profile
	err := Models.GetSearch(&profile, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

//Updateprofile ... Update the profile information
func UpdateProfile(c *gin.Context) {
	var profile Models.Profile
	id := c.Params.ByName("id")
	err := Models.GetProfileByID(&profile, id)
	if err != nil {
		c.JSON(http.StatusNotFound, profile)
	}
	c.BindJSON(&profile)
	err = Models.UpdateProfile(&profile, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

//Deleteprofile ... Delete the profile
func DeleteProfile(c *gin.Context) {
	var profile Models.Profile
	id := c.Params.ByName("id")
	err := Models.DeleteProfile(&profile, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
