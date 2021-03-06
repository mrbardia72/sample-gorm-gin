package routes

import (
	"github.com/mrbardia72/sample-gorm-gin/controllers/user"
	"github.com/mrbardia72/sample-gorm-gin/controllers/post"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//routers user
	grp1 := r.Group("/v1/api/user")
	{
		grp1.GET("/count", user.CountUser)
		grp1.GET("/read", user.GetUsers)
		grp1.POST("/create", user.CreateUser)
		grp1.GET("/search/:name", user.GetSearchUser)
		grp1.PUT("/update/:id", user.UpdateUser)
		grp1.DELETE("/delete/:id", user.DeleteUser)
	}

	// routes posts
	grp2 := r.Group("/v2/api/post")
	{
		grp2.GET("/count", post.CountPost)
		grp2.GET("/read", post.Getposts)
		grp2.POST("/create",post.CreatePost)
		grp2.PUT("/update/:id", post.UpdatePost)
		grp2.DELETE("/delete/:id", post.DeletePost)
	}
	return r
}
