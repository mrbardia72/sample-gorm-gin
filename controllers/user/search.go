package user

import (
	// "fmt"
	repo_user "github.com/mrbardia72/sample-gorm-gin/repository/user"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//GetSearchUser ... Get the user by name
func GetSearchUser(c *gin.Context) {
	name := c.Params.ByName("name") 
	var userx models.User
	err := repo_user.GetSearchUser(&userx, name)
	if err != nil {
		errorx:=helpers.Errorx{Msgx:"There is no such recordxxx",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, userx)
	}
}