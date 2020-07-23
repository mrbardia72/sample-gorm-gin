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
		return err
	} 
	info:=" Successfully read all post "
	helpers.LogApi(info)
	return nil
}