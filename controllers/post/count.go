package post

import (
	"net/http"
	"github.com/mrbardia72/sample-gorm-gin/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"

)
func CountPost(c *gin.Context) {
 
	var result int64
	config.DB.Table("posts").Count(&result)
	c.JSON(http.StatusOK, result)
}