package post

import (
	repo_post "github.com/mrbardia72/sample-gorm-gin/repository/post"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)

//Updatepost ... Update the post information
func UpdatePost(c *gin.Context) {
	var postx models.Post
	id := c.Params.ByName("id")
	err := repo_post.GetSearchPost(&postx, id) 
	if err != nil {
		c.JSON(http.StatusNotFound, postx)
	}
	c.BindJSON(&postx)
	err = repo_post.UpdatePost(&postx, id)
	if err != nil {
		errorx:=helpers.Errorx{Msgx:"not update a record",Codex:"404"}
		c.JSON(http.StatusNotFound,errorx)
	} else {
		c.JSON(http.StatusOK, postx)
	}
}