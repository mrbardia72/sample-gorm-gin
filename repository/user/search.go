package user
 
import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
) 

func GetSearchUser(user *models.User, name string) (err error) {

	if err = config.DB.Where("name = ?", name).Find(&user).Error; err != nil {
		
		info:="There is no such record"
		helpers.LogApi(info)
		return 

	} else { 

		config.DB.Preload("Posts").Find(&user)
		info:=" Successfully found  search User "
		helpers.LogApi(info)
		return
	}
}