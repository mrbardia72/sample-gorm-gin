package post
 
import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
) 

func GetSearchPost(post *models.Post, title string) (err error) {

	if err = config.DB.Where("title = ?", title).Find(&post).Error; err != nil {
		
		info:="There is no such record"
		helpers.LogApi(info)
		return 

	} else { 

		// config.DB.Preload("Posts").Find(&post)
		info:=" Successfully found  search post "
		helpers.LogApi(info)
		return
	}
}