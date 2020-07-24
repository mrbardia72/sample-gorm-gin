package post

import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
)
	
func UpdatePost(post *models.Post, id string) (err error) {
	fmt.Println(post)
	config.DB.Save(post)
	return nil
}