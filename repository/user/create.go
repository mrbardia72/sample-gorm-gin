package user

import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
   
func CreateUser(user *models.User) (err error) {

	if err = config.DB.Create(user).Error; err != nil {

		info:=" error, create user"
		helpers.LogApi(info)
		return err
	}

	info:=" ok, create user "
	helpers.LogApi(info)
	return nil
}