package user

import (
	// "fmt"
	repo_user "github.com/mrbardia72/sample-gorm-gin/repository/user"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var userx models.User
	id := c.Params.ByName("id")
	err := repo_user.DeleteUser(&userx, id)

	if err != nil {
		errorx:=helpers.Errorx{Msgx:"not delete a record",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
	
}