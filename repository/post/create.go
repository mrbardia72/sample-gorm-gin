package post

import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
   
func CreatePost(post *models.Post) (err error) {

	if err = config.DB.Create(post).Error; err != nil {

		info:=" error, create post"
		helpers.LogApi(info)
		return err
	}

	info:=" ok, create post "
	helpers.LogApi(info)
	return nil
}