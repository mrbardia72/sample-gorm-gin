package user
 
import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
)
	
func UpdateUser(user *models.User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}