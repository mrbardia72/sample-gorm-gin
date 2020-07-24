package post

import (
	// "fmt"
	repo_post "github.com/mrbardia72/sample-gorm-gin/repository/post"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//Createpost ... Create post
func CreatePost(c *gin.Context) {
	var postx models.Post

	if err := c.ShouldBindJSON(&postx); err != nil {

		errorx:=helpers.Errorx{Msgx:"Status Bad Request",Codex:"400"}
		c.JSON(http.StatusBadRequest, errorx)
		return
	}

	c.BindJSON(&postx)
	err := repo_post.CreatePost(&postx)

	if err != nil {

		errorx:=helpers.Errorx{Msgx:"error create to db",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
		
	} else {

		c.JSON(http.StatusOK, postx)
	}
}