package routes

import (
	"github.com/mrbardia72/sample-gorm-gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//routers profile(users)
	grp1 := r.Group("/v1/api/profile")
	{
		grp1.GET("/read", controllers.GetProfiles)
		grp1.POST("/create", controllers.CreateProfile)
		grp1.GET("/search/:name", controllers.GetSearch)
		grp1.PUT("/update/:id", controllers.UpdateProfile)
		grp1.DELETE("/delete/:id", controllers.DeleteProfile)
	}

	// routes posts
	// grp2 := r.Group("/v1/api/post")
	// {
		
	// }

	return r
}
