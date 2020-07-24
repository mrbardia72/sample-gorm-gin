package user

import (
	// "fmt"
	repo_user "github.com/mrbardia72/sample-gorm-gin/repository/user"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

  
//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var userx []models.User
	err := repo_user.GetAllUsers(&userx)

	if err != nil {
		errorx:=helpers.Errorx{Msgx:"error not get all users becaus empty table user",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, userx)
	}

}