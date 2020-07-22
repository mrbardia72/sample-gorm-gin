package routes

import (
	"github.com/mrbardia72/sample-gorm-gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//routers user
	grp1 := r.Group("/v1/api/user")
	{
		grp1.GET("/read", controllers.GetUsers)
		grp1.POST("/create", controllers.CreateUser)
		grp1.GET("/search/:name", controllers.GetSearchUser)
		grp1.PUT("/update/:id", controllers.UpdateUser)
		grp1.DELETE("/delete/:id", controllers.DeleteUser)
	}

	// routes posts
	grp2 := r.Group("/v2/api/post")
	{
		grp2.GET("/read", controllers.Getposts)
	}
	return r
}
