package user
 
import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers" 
)

func GetUserByID(user *models.User, id string) (err error) {
	
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {

		info:=" error, get user by id "
		helpers.LogApi(info)
	 	return err
	}

	info:=" ok, get  user by id"
	helpers.LogApi(info)
	return nil
   }