package user

import (
	// "fmt"
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
  
//GetAllUsers Fetch all user data
func GetAllUsers(user *[]models.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		info:=" ok, get all users "
		helpers.LogApi(info)
		return err
	}
	info:=" error, get all users "
	helpers.LogApi(info)
	return nil
}