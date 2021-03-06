package post

import (
	repo_post "github.com/mrbardia72/sample-gorm-gin/repository/post"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
 
func Getposts(c *gin.Context) {
	var postx []models.Post
	err := repo_post.GetAllposts(&postx)

	if err != nil {
		errorx:=helpers.Errorx{Msgx:"error = not get all posts = because = empty table user",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, postx)
	}
}