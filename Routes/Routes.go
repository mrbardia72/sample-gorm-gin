package Routes

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/v1/api/profile")
	{
		grp1.GET("/read", Controllers.GetProfiles)
		grp1.POST("/create", Controllers.CreateProfile)
		grp1.GET("/search/:name", Controllers.GetSearch)
		grp1.PUT("/update/:id", Controllers.UpdateProfile)
		grp1.DELETE("/delete/:id", Controllers.DeleteProfile)
	}
	return r
}
