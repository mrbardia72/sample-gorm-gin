package post

import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
)

func DeletePost(post *models.Post, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(post)
	return nil
}