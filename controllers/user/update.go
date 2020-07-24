package user

import (
	// "fmt"
	repo_user "github.com/mrbardia72/sample-gorm-gin/repository/user"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var userx models.User
	id := c.Params.ByName("id")
	err := repo_user.GetSearchUser(&userx, id) 
	if err != nil {
		c.JSON(http.StatusNotFound, userx)
	}
	c.BindJSON(&userx)
	err = repo_user.UpdateUser(&userx, id)
	if err != nil {
		errorx:=helpers.Errorx{Msgx:"not update a record",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, userx)
	}
}