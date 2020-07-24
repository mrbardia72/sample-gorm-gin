package post

import (
	// "fmt"
	repo_post "github.com/mrbardia72/sample-gorm-gin/repository/post"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//DeletePost ... Delete the Post
func DeletePost(c *gin.Context) {
	var postx models.Post
	id := c.Params.ByName("id")
	err := repo_post.DeletePost(&postx, id)

	if err != nil {
		errorx:=helpers.Errorx{Msgx:"not delete a record",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
	
}