package user

import (
	// "fmt"
	repo_user "github.com/mrbardia72/sample-gorm-gin/repository/user"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//CreateUser ... Create user
func CreateUser(c *gin.Context) {
	var userx models.User

	if err := c.ShouldBindJSON(&userx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&userx)
	err := repo_user.CreateUser(&userx)

	if err != nil {
		errorx:=helpers.Errorx{Msgx:"error create to db",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, userx)
	}
}