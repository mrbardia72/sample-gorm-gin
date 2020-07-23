package user

import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteUser(user *models.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}