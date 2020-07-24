package post

import (
	// "fmt"
	repo_post "github.com/mrbardia72/sample-gorm-gin/repository/post"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//GetSearchpost ... Get the post by name
func GetSearchPost(c *gin.Context) {
	name := c.Params.ByName("title") 
	var postx models.Post
	err := repo_post.GetSearchPost(&postx, title)
	if err != nil {
		errorx:=helpers.Errorx{Msgx:"There is no such recordxxx",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, postx)
	}
}