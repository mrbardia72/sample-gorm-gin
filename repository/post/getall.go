package post

import (
	// "fmt"
	"github.com/mrbardia72/sample-gorm-gin/config"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"github.com/mrbardia72/sample-gorm-gin/helpers"
	_ "github.com/go-sql-driver/mysql"
)  

//GetAllposts Fetch all user data
func GetAllposts(post *[]models.Post) (err error) {
	if err = config.DB.Find(post).Error; err != nil {
		info:=" error, get all posts "
		helpers.LogApi(info)
		return err
	} 
	info:=" ok, get all posts "
	helpers.LogApi(info)
	return nil
}